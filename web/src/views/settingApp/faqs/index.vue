<template>
  <ele-page>
    <ele-card
      :body-style="{ paddingBottom: '4px' }"
      :style="{
        minHeight: '380px',
        marginBottom: '0px'
      }"
    >
      <template #header>
        <div class="page-header">
          <div class="header-icon">
            <el-icon :size="24" color="#409eff">
              <QuestionFilled />
            </el-icon>
          </div>
          <div class="header-content">
            <h1 class="header-title">{{ t('FAQs.title') }}</h1>
            <p class="header-description">{{ t('FAQs.description') }}</p>
          </div>
        </div>
      </template>

      <el-tabs v-model="activeTab" class="main-tabs">
        <el-tab-pane :label="t('FAQs.faqTab')" name="faq">
          <div class="tab-content">
            <div class="toolbar-section">
              <el-button type="primary" @click="handleAddFaq" class="ele-btn-icon">
                <el-icon><Plus /></el-icon>
                {{ t('common.add') }} FAQ
              </el-button>
            </div>
            <div v-if="showFaqForm" class="form-section">
              <ele-card header="FAQ Form" shadow="hover">
                <el-form :model="faqForm" label-width="120px" @submit.prevent="saveFaq">
                  <el-form-item label="Question">
                    <el-input v-model="faqForm.question" placeholder="Enter FAQ question" />
                  </el-form-item>
                  <el-form-item label="Answer">
                    <el-input v-model="faqForm.answer" type="textarea" :rows="4" placeholder="Enter FAQ answer" />
                  </el-form-item>
                  <el-form-item label="Category">
                    <el-select v-model="faqForm.categoryId" placeholder="Select category" style="width: 100%">
                      <el-option v-for="category in categories" :key="category.id" :label="category.name" :value="category.id" />
                    </el-select>
                  </el-form-item>
                  <el-form-item>
                    <el-button type="primary" @click="saveFaq">{{ t('common.save') }}</el-button>
                    <el-button @click="cancelFaq">{{ t('common.cancel') }}</el-button>
                  </el-form-item>
                </el-form>
              </ele-card>
            </div>
            <div v-if="loading" class="loading-section">
              <el-skeleton :rows="3" animated />
            </div>
            <div v-else-if="error" class="error-section">
              <el-alert type="error" :title="error" show-icon />
            </div>
            <div v-else class="content-section">
              <div v-for="faq in faqs" :key="faq.id" class="faq-item-card">
                <div class="faq-content">
                  <div class="faq-question">{{ faq.question }}</div>
                  <div class="faq-answer">{{ faq.answer }}</div>
                  <div class="faq-category">
                    <el-tag v-if="faq.category" size="small">{{ faq.category.name }}</el-tag>
                  </div>
                </div>
                <div class="faq-actions">
                  <el-button size="small" type="primary" text @click="editFaq(faq)">
                    <el-icon><Edit /></el-icon>
                  </el-button>
                  <el-button size="small" type="danger" text @click="deleteFaq(faq.id)">
                    <el-icon><Delete /></el-icon>
                  </el-button>
                </div>
              </div>
            </div>
          </div>
        </el-tab-pane>
        <el-tab-pane :label="t('FAQs.categoriesTab')" name="categories">
          <div class="tab-content">
            <div class="toolbar-section">
              <el-button type="primary" @click="handleAddCategory" class="ele-btn-icon">
                <el-icon><Plus /></el-icon>
                {{ t('common.add') }} Category
              </el-button>
            </div>
            <div v-if="showCategoryForm" class="form-section">
              <ele-card header="Category Form" shadow="hover">
                <el-form :model="categoryForm" label-width="120px" @submit.prevent="saveCategory">
                  <el-form-item label="Name">
                    <el-input v-model="categoryForm.name" placeholder="Enter category name" />
                  </el-form-item>
                  <el-form-item label="Description">
                    <el-input v-model="categoryForm.description" type="textarea" :rows="3" placeholder="Enter category description" />
                  </el-form-item>
                  <el-form-item>
                    <el-button type="primary" @click="saveCategory">{{ t('common.save') }}</el-button>
                    <el-button @click="cancelCategory">{{ t('common.cancel') }}</el-button>
                  </el-form-item>
                </el-form>
              </ele-card>
            </div>
            <div v-if="categoriesLoading" class="loading-section">
              <el-skeleton :rows="3" animated />
            </div>
            <div v-else-if="categoriesError" class="error-section">
              <el-alert type="error" :title="categoriesError" show-icon />
            </div>
            <div v-else class="content-section">
              <div v-for="category in categories" :key="category.id" class="category-item-card">
                <div class="category-content">
                  <div class="category-name">{{ category.name }}</div>
                  <div class="category-description">{{ category.description }}</div>
                  <div class="category-stats">
                    <el-tag size="small">{{ category.faqCount || 0 }} FAQs</el-tag>
                  </div>
                </div>
                <div class="category-actions">
                  <el-button size="small" type="primary" text @click="editCategory(category)">
                    <el-icon><Edit /></el-icon>
                  </el-button>
                  <el-button size="small" type="danger" text @click="deleteCategory(category.id)">
                    <el-icon><Delete /></el-icon>
                  </el-button>
                </div>
              </div>
            </div>
          </div>
        </el-tab-pane>
      </el-tabs>
    </ele-card>
  </ele-page>
</template>

<script setup>
  import { ElMessage, ElMessageBox } from 'element-plus';
  import { onMounted, reactive, ref } from 'vue';
  import { useI18n } from 'vue-i18n';

  const { t } = useI18n();

  const activeTab = ref('faq');

  const faqs = ref([]);
  const loading = ref(false);
  const error = ref('');
  const showFaqForm = ref(false);
  const faqForm = reactive({
    id: null,
    question: '',
    answer: '',
    categoryId: null
  });

  const categories = ref([]);
  const categoriesLoading = ref(false);
  const categoriesError = ref('');
  const showCategoryForm = ref(false);
  const categoryForm = reactive({
    id: null,
    name: '',
    description: ''
  });

  const fetchFaqs = async () => {
    loading.value = true;
    error.value = '';
    try {
      await new Promise((resolve) => setTimeout(resolve, 1000));
      faqs.value = [
        {
          id: 1,
          question: 'How to reset password?',
          answer: 'You can reset your password by clicking on the "Forgot Password" link on the login page.',
          categoryId: 1,
          category: { id: 1, name: 'Account' }
        },
        {
          id: 2,
          question: 'How to contact support?',
          answer: 'You can contact our support team through the help center or by sending an email to support@example.com.',
          categoryId: 2,
          category: { id: 2, name: 'Support' }
        }
      ];
    } catch (err) {
      error.value = 'Failed to load FAQs';
    } finally {
      loading.value = false;
    }
  };

  const fetchCategories = async () => {
    categoriesLoading.value = true;
    categoriesError.value = '';
    try {
      await new Promise((resolve) => setTimeout(resolve, 800));
      categories.value = [
        { id: 1, name: 'Account', description: 'Account related questions', faqCount: 5 },
        { id: 2, name: 'Support', description: 'Support and help questions', faqCount: 3 },
        { id: 3, name: 'Billing', description: 'Billing and payment questions', faqCount: 2 }
      ];
    } catch (err) {
      categoriesError.value = 'Failed to load categories';
    } finally {
      categoriesLoading.value = false;
    }
  };

  const handleAddFaq = () => {
    resetFaqForm();
    showFaqForm.value = true;
  };

  const editFaq = (faq) => {
    faqForm.id = faq.id;
    faqForm.question = faq.question;
    faqForm.answer = faq.answer;
    faqForm.categoryId = faq.categoryId;
    showFaqForm.value = true;
  };

  const saveFaq = async () => {
    try {
      await new Promise((resolve) => setTimeout(resolve, 500));
      if (faqForm.id) {
        const index = faqs.value.findIndex((f) => f.id === faqForm.id);
        if (index !== -1) {
          const category = categories.value.find((c) => c.id === faqForm.categoryId);
          faqs.value[index] = {
            ...faqForm,
            category
          };
        }
        ElMessage.success('FAQ updated successfully');
      } else {
        const category = categories.value.find((c) => c.id === faqForm.categoryId);
        const newFaq = {
          id: Date.now(),
          ...faqForm,
          category
        };
        faqs.value.push(newFaq);
        ElMessage.success('FAQ added successfully');
      }

      showFaqForm.value = false;
      resetFaqForm();
    } catch (err) {
      ElMessage.error('Failed to save FAQ');
    }
  };

  const deleteFaq = async (id) => {
    try {
      await ElMessageBox.confirm('Are you sure you want to delete this FAQ?', 'Confirm Delete', {
        type: 'warning'
      });
      await new Promise((resolve) => setTimeout(resolve, 300));
      faqs.value = faqs.value.filter((f) => f.id !== id);
      ElMessage.success('FAQ deleted successfully');
    } catch (err) {
      if (err !== 'cancel') {
        ElMessage.error('Failed to delete FAQ');
      }
    }
  };

  const cancelFaq = () => {
    showFaqForm.value = false;
    resetFaqForm();
  };

  const resetFaqForm = () => {
    faqForm.id = null;
    faqForm.question = '';
    faqForm.answer = '';
    faqForm.categoryId = null;
  };

  const handleAddCategory = () => {
    resetCategoryForm();
    showCategoryForm.value = true;
  };

  const editCategory = (category) => {
    categoryForm.id = category.id;
    categoryForm.name = category.name;
    categoryForm.description = category.description;
    showCategoryForm.value = true;
  };

  const saveCategory = async () => {
    try {
      await new Promise((resolve) => setTimeout(resolve, 500));
      if (categoryForm.id) {
        const index = categories.value.findIndex((c) => c.id === categoryForm.id);
        if (index !== -1) {
          categories.value[index] = { ...categoryForm, faqCount: categories.value[index].faqCount };
        }
        ElMessage.success('Category updated successfully');
      } else {
        const newCategory = {
          id: Date.now(),
          ...categoryForm,
          faqCount: 0
        };
        categories.value.push(newCategory);
        ElMessage.success('Category added successfully');
      }

      showCategoryForm.value = false;
      resetCategoryForm();
    } catch (err) {
      ElMessage.error('Failed to save category');
    }
  };

  const deleteCategory = async (id) => {
    try {
      await ElMessageBox.confirm('Are you sure you want to delete this category?', 'Confirm Delete', {
        type: 'warning'
      });
      await new Promise((resolve) => setTimeout(resolve, 300));

      categories.value = categories.value.filter((c) => c.id !== id);
      ElMessage.success('Category deleted successfully');
    } catch (err) {
      if (err !== 'cancel') {
        ElMessage.error('Failed to delete category');
      }
    }
  };

  const cancelCategory = () => {
    showCategoryForm.value = false;
    resetCategoryForm();
  };

  const resetCategoryForm = () => {
    categoryForm.id = null;
    categoryForm.name = '';
    categoryForm.description = '';
  };

  onMounted(() => {
    fetchFaqs();
    fetchCategories();
  });
</script>

<style lang="scss" scoped>
  .page-header {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 8px 0;

    .header-icon {
      flex-shrink: 0;
      width: 48px;
      height: 48px;
      border-radius: 8px;
      overflow: hidden;
      border: 1px solid #e4e7ed;
      display: flex;
      align-items: center;
      justify-content: center;
      background: #f5f7fa;
    }

    .header-content {
      flex: 1;
      min-width: 0;

      .header-title {
        font-size: 18px;
        font-weight: 600;
        color: #303133;
        margin-bottom: 4px;
        line-height: 1.4;
      }

      .header-description {
        font-size: 14px;
        color: #909399;
        line-height: 1.3;
      }
    }
  }

  .main-tabs {
    :deep(.el-tabs__header) {
      margin-bottom: 16px;
    }

    :deep(.el-tabs__content) {
      padding: 0;
    }
  }

  .tab-content {
    padding: 16px 0;
  }

  .toolbar-section {
    margin-bottom: 16px;
    display: flex;
    justify-content: flex-start;
    align-items: center;
  }

  .form-section {
    margin-bottom: 24px;
  }

  .loading-section,
  .error-section {
    text-align: center;
    padding: 32px 0;
  }

  .content-section {
    margin-top: 16px;
  }

  .faq-item-card {
    display: flex;
    align-items: flex-start;
    gap: 12px;
    padding: 16px;
    border: 1px solid #e4e7ed;
    border-radius: 8px;
    margin-bottom: 12px;
    background: #fff;
    transition: all 0.3s ease;

    &:hover {
      border-color: #409eff;
      box-shadow: 0 2px 8px rgba(64, 158, 255, 0.1);
    }

    .faq-content {
      flex: 1;
      min-width: 0;

      .faq-question {
        font-size: 14px;
        font-weight: 600;
        color: #303133;
        margin-bottom: 8px;
        line-height: 1.4;
      }

      .faq-answer {
        font-size: 13px;
        color: #606266;
        line-height: 1.5;
        margin-bottom: 8px;
      }

      .faq-category {
        font-size: 12px;
        color: #909399;
      }
    }

    .faq-actions {
      display: flex;
      gap: 8px;
      flex-shrink: 0;
    }
  }

  .category-item-card {
    display: flex;
    align-items: flex-start;
    gap: 12px;
    padding: 16px;
    border: 1px solid #e4e7ed;
    border-radius: 8px;
    margin-bottom: 12px;
    background: #fff;
    transition: all 0.3s ease;

    &:hover {
      border-color: #67c23a;
      box-shadow: 0 2px 8px rgba(103, 194, 58, 0.1);
    }

    .category-content {
      flex: 1;
      min-width: 0;

      .category-name {
        font-size: 14px;
        font-weight: 600;
        color: #303133;
        margin-bottom: 4px;
        line-height: 1.4;
      }

      .category-description {
        font-size: 13px;
        color: #606266;
        line-height: 1.5;
      }
    }

    .category-actions {
      display: flex;
      gap: 8px;
      flex-shrink: 0;
    }
  }
</style>
