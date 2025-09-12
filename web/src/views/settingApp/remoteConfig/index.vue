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
              <Setting />
            </el-icon>
          </div>
          <div class="header-content">
            <h1 class="header-title">{{ $t('remoteConfig.title') }}</h1>
            <p class="header-description">{{ $t('remoteConfig.description') }}</p>
          </div>
        </div>
      </template>
      <el-tabs v-model="activeTab" class="main-tabs">
        <el-tab-pane :label="$t('remoteConfig.configTab')" name="config">
          <div class="tab-content">
            <div class="toolbar-section">
              <el-button
                type="primary"
                @click="handleConfigSubmit"
                :loading="isUpdatingConfig"
              >
                {{ isUpdatingConfig ? $t('remoteConfig.updating') : $t('remoteConfig.updateConfiguration') }}
              </el-button>
            </div>
            <div v-if="isLoadingConfig" class="loading-section">
              <el-skeleton :rows="6" animated />
            </div>
            <div v-else-if="configError" class="error-section">
              <el-alert
                :title="$t('common.error')"
                :description="configError"
                type="error"
                show-icon
                :closable="false"
              />
            </div>
            <div v-else class="content-section">
              <div class="status-section">
                <el-tag
                  :type="remoteConfigForm.isEnabled ? 'success' : 'info'"
                  size="large"
                >
                  {{ remoteConfigForm.isEnabled ? $t('remoteConfig.active') : $t('remoteConfig.inactive') }}
                </el-tag>
                <span v-if="lastUpdated" class="last-updated">
                  {{ $t('remoteConfig.lastUpdated') }}: {{ formatDate(lastUpdated) }}
                </span>
              </div>
              <ele-card class="config-card">
                <template #header>
                  <div class="card-header">
                    <div class="card-title">
                      <el-icon><Message /></el-icon>
                      <h3>{{ $t('remoteConfig.configurationSettings') }}</h3>
                    </div>
                    <p>{{ $t('remoteConfig.configurationSettingsDescription') }}</p>
                  </div>
                </template>
                <el-form
                  ref="configFormRef"
                  :model="remoteConfigForm"
                  :rules="configRules"
                  label-width="120px"
                  class="config-form"
                  @submit.prevent="handleConfigSubmit"
                >
                <el-form-item :label="$t('remoteConfig.langField')" prop="lang">
                  <el-select
                    v-model="remoteConfigForm.lang"
                    :placeholder="$t('remoteConfig.langPlaceholder')"
                    style="width: 100%"
                    @change="loadRemoteConfig"
                  >
                    <el-option :label="$t('nav.english')" value="EN" />
                    <el-option :label="$t('nav.chinese')" value="CN" />
                    <el-option :label="$t('nav.thai')" value="TAI" />
                  </el-select>
                  <div class="form-description">{{ $t('remoteConfig.langDescription') }}</div>
                </el-form-item>
                <el-divider />
                <el-form-item :label="$t('remoteConfig.titleField')" prop="title">
                  <el-input
                    v-model="remoteConfigForm.title"
                    :placeholder="$t('remoteConfig.titlePlaceholder')"
                  />
                  <div class="form-description">{{ $t('remoteConfig.titleDescription') }}</div>
                </el-form-item>
                <el-divider />
                <el-form-item :label="$t('remoteConfig.messageField')" prop="message">
                  <el-input
                    v-model="remoteConfigForm.message"
                    type="textarea"
                    :rows="4"
                    :placeholder="$t('remoteConfig.messagePlaceholder')"
                    resize="none"
                  />
                  <div class="form-description">{{ $t('remoteConfig.messageDescription') }}</div>
                </el-form-item>
                <el-divider />
                <el-form-item :label="$t('remoteConfig.buttonTextField')" prop="buttonText">
                  <el-input
                    v-model="remoteConfigForm.buttonText"
                    :placeholder="$t('remoteConfig.buttonTextPlaceholder')"
                  />
                  <div class="form-description">{{ $t('remoteConfig.buttonTextDescription') }}</div>
                </el-form-item>
                <el-divider />
                <el-form-item :label="$t('remoteConfig.enable')">
                  <div class="switch-container">
                    <div class="switch-content">
                      <div class="switch-info">
                        <h4>{{ $t('remoteConfig.enableConfiguration') }}</h4>
                        <p>{{ $t('remoteConfig.enableConfigurationDescription') }}</p>
                      </div>
                      <el-switch
                        v-model="remoteConfigForm.isEnabled"
                        size="large"
                      />
                    </div>
                  </div>
                </el-form-item>
                <el-form-item class="submit-section">
                  <el-button
                    type="primary"
                    :loading="isUpdatingConfig"
                    @click="handleConfigSubmit"
                    size="large"
                  >
                    {{ isUpdatingConfig ? $t('remoteConfig.updating') : $t('remoteConfig.updateConfiguration') }}
                  </el-button>
                </el-form-item>
                </el-form>
              </ele-card>
            </div>
          </div>
        </el-tab-pane>

        <!-- FAQs Tab -->
        <el-tab-pane :label="$t('FAQs.title')" name="faqs">
          <div class="tab-content">
            <div class="toolbar-section">
              <el-button
                type="primary"
                @click="handleAddFaq"
                size="large"
              >
                <el-icon><Plus /></el-icon>
                {{ $t('FAQs.addFaq') }}
              </el-button>
            </div>

            <el-tabs v-model="activeFaqTab" class="faq-tabs">
              <el-tab-pane :label="$t('FAQs.faqTab')" name="faq">
                <div class="faq-content">
                  <div v-if="isLoadingFaqs" class="loading-section">
                    <el-skeleton :rows="4" animated v-for="i in 3" :key="i" />
                  </div>
                  <div v-else-if="faqsError" class="error-section">
                    <el-alert
                      :title="$t('common.error')"
                      :description="faqsError"
                      type="error"
                      show-icon
                      :closable="false"
                    />
                  </div>
                  <div v-else-if="localFaqs.length === 0" class="empty-state">
                    <ele-card class="empty-card">
                      <div class="empty-content">
                        <div class="empty-icon">
                          <el-icon><QuestionFilled /></el-icon>
                        </div>
                        <h3>{{ $t('FAQs.noFaqsYet') }}</h3>
                        <p>{{ $t('FAQs.noFaqsDescription') }}</p>
                        <el-button type="primary" @click="handleAddFaq" size="large">
                          <el-icon><Plus /></el-icon>
                          {{ $t('FAQs.addFaq') }}
                        </el-button>
                      </div>
                    </ele-card>
                  </div>
                  <div v-else class="content-section">
                    <div class="faqs-list">
                    <ele-card
                      v-for="(faq, index) in localFaqs"
                      :key="faq.id || `temp-${index}`"
                      class="faq-card"
                    >
                      <template #header>
                        <div class="faq-header">
                          <div class="faq-info">
                            <el-tag v-if="!faq.id" type="warning">{{ $t('FAQs.new') }}</el-tag>
                            <el-tag v-else-if="editingFaqIndex === index" type="primary">{{ $t('FAQs.editing') }}</el-tag>
                          </div>
                          <div class="faq-actions" v-if="editingFaqIndex !== index">
                            <el-button
                              type="primary"
                              text
                              @click="handleEditFaq(index)"
                            >
                              <el-icon><Edit /></el-icon>
                            </el-button>
                            <el-button
                              type="danger"
                              text
                              @click="handleDeleteFaq(index)"
                            >
                              <el-icon><Delete /></el-icon>
                            </el-button>
                          </div>
                        </div>
                      </template>
                      <div v-if="editingFaqIndex === index" class="faq-edit">
                        <el-form
                          ref="faqFormRef"
                          :model="editingFaq"
                          :rules="faqRules"
                          label-width="100px"
                        >
                          <el-form-item :label="$t('FAQs.category')" prop="categoryId">
                            <el-select
                              v-model="editingFaq.categoryId"
                              :placeholder="$t('FAQs.selectCategory')"
                              style="width: 100%"
                            >
                              <el-option
                                v-for="category in categories"
                                :key="category.id"
                                :label="category.name"
                                :value="category.id"
                              />
                            </el-select>
                          </el-form-item>
                          <el-form-item :label="$t('FAQs.question')" prop="question">
                            <el-input
                              v-model="editingFaq.question"
                              :placeholder="$t('FAQs.questionPlaceholder')"
                            />
                          </el-form-item>
                          <el-form-item :label="$t('FAQs.answer')" prop="answer">
                            <el-input
                              v-model="editingFaq.answer"
                              type="textarea"
                              :rows="4"
                              :placeholder="$t('FAQs.answerPlaceholder')"
                            />
                          </el-form-item>
                          <el-form-item>
                            <el-button
                              type="primary"
                              @click="handleSaveFaq(index)"
                              :loading="isSavingFaq"
                            >
                              <el-icon><Check /></el-icon>
                              {{ isSavingFaq ? $t('FAQs.saving') : $t('FAQs.save') }}
                            </el-button>
                            <el-button @click="handleCancelEditFaq">
                              <el-icon><Close /></el-icon>
                              {{ $t('FAQs.cancel') }}
                            </el-button>
                          </el-form-item>
                        </el-form>
                      </div>
                      <div v-else class="faq-view">
                        <div class="faq-question">
                          <h4>{{ faq.question || $t('FAQs.untitledQuestion') }}</h4>
                        </div>
                        <div class="faq-answer">
                          <p>{{ faq.answer || $t('FAQs.noAnswer') }}</p>
                        </div>
                        <div class="faq-category" v-if="faq.categoryId">
                          <el-tag size="small">{{ getCategoryName(faq.categoryId) }}</el-tag>
                        </div>
                      </div>
                    </ele-card>
                  </div>
                </div>
              </div>
              </el-tab-pane>
              <el-tab-pane :label="$t('FAQs.categoriesTab')" name="categories">
                <div class="categories-content">
                  <div class="add-section">
                    <el-button
                      type="primary"
                      @click="handleAddCategory"
                      size="large"
                    >
                      <el-icon><Plus /></el-icon>
                      {{ $t('FAQs.addCategory') }}
                    </el-button>
                  </div>
                  <div v-if="isLoadingCategories" class="loading-section">
                    <el-skeleton :rows="3" animated v-for="i in 3" :key="i" />
                  </div>
                  <div v-else-if="categoriesError" class="error-section">
                    <el-alert
                      :title="$t('common.error')"
                      :description="categoriesError"
                      type="error"
                      show-icon
                      :closable="false"
                    />
                  </div>
                  <div v-else-if="localCategories.length === 0" class="empty-state">
                    <ele-card class="empty-card">
                      <div class="empty-content">
                        <div class="empty-icon">
                          <el-icon><FolderOpened /></el-icon>
                        </div>
                        <h3>{{ $t('FAQs.noCategoriesYet') }}</h3>
                        <p>{{ $t('FAQs.noCategoriesDescription') }}</p>
                        <el-button type="primary" @click="handleAddCategory" size="large">
                          <el-icon><Plus /></el-icon>
                          {{ $t('FAQs.addCategory') }}
                        </el-button>
                      </div>
                    </ele-card>
                  </div>
                  <div v-else class="content-section">
                    <div class="categories-grid">
                    <ele-card
                      v-for="(category, index) in allCategories"
                      :key="category.id || `temp-${index}`"
                      class="category-card"
                    >
                      <template #header>
                        <div class="category-header">
                          <div class="category-info">
                            <el-tag v-if="!category.id" type="warning">{{ $t('FAQs.new') }}</el-tag>
                            <el-tag v-else-if="editingCategoryId === category.id" type="primary">{{ $t('FAQs.editing') }}</el-tag>
                          </div>
                          <div class="category-actions" v-if="editingCategoryId !== category.id">
                            <el-button
                              type="primary"
                              text
                              @click="handleEditCategory(category)"
                            >
                              <el-icon><Edit /></el-icon>
                            </el-button>
                            <el-button
                              type="danger"
                              text
                              @click="handleDeleteCategory(category)"
                            >
                              <el-icon><Delete /></el-icon>
                            </el-button>
                          </div>
                        </div>
                      </template>
                      <div v-if="editingCategoryId === category.id" class="category-edit">
                        <el-form
                          ref="categoryFormRef"
                          :model="editingCategory"
                          :rules="categoryRules"
                          label-width="100px"
                        >
                          <el-form-item :label="$t('FAQs.categoryName')" prop="name">
                            <el-input
                              v-model="editingCategory.name"
                              :placeholder="$t('FAQs.categoryNamePlaceholder')"
                            />
                          </el-form-item>
                          <el-form-item>
                            <el-button
                              type="primary"
                              @click="handleSaveCategory"
                              :loading="isSavingCategory"
                            >
                              <el-icon><Check /></el-icon>
                              {{ isSavingCategory ? $t('FAQs.saving') : $t('FAQs.save') }}
                            </el-button>
                            <el-button @click="handleCancelEditCategory">
                              <el-icon><Close /></el-icon>
                              {{ $t('FAQs.cancel') }}
                            </el-button>
                          </el-form-item>
                        </el-form>
                      </div>
                      <div v-else class="category-view">
                        <h4>{{ category.name || $t('FAQs.untitledCategory') }}</h4>
                        <p class="category-description">{{ $t('FAQs.categoryDescription') }}</p>
                      </div>
                    </ele-card>
                    </div>
                  </div>
                </div>
              </el-tab-pane>
            </el-tabs>
          </div>
        </el-tab-pane>
      </el-tabs>
    </ele-card>
  </ele-page>
</template>

<script>
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Setting,
  Message,
  QuestionFilled,
  Plus,
  Edit,
  Delete,
  Check,
  Close,
  FolderOpened
} from '@element-plus/icons-vue'

export default {
  name: 'RemoteConfigSettings',
  components: {
    Setting,
    Message,
    QuestionFilled,
    Plus,
    Edit,
    Delete,
    Check,
    Close,
    FolderOpened
  },
  setup() {

    const activeTab = ref('config')
    const activeFaqTab = ref('faq')


    const isLoadingConfig = ref(false)
    const isUpdatingConfig = ref(false)
    const configError = ref(null)
    const lastUpdated = ref(null)
    const configFormRef = ref(null)

    const remoteConfigForm = reactive({
      title: '',
      message: '',
      buttonText: '',
      isEnabled: false,
      lang: 'EN'
    })

    const isLoadingFaqs = ref(false)
    const isLoadingCategories = ref(false)
    const isSavingFaq = ref(false)
    const isSavingCategory = ref(false)
    const editingFaqIndex = ref(null)
    const editingCategoryId = ref(null)
    const faqFormRef = ref(null)
    const categoryFormRef = ref(null)

    const localFaqs = ref([])
    const localCategories = ref([])
    const newCategories = ref([])
    const categories = ref([
      { id: '1', name: 'General' },
      { id: '2', name: 'Technical' },
      { id: '3', name: 'Billing' }
    ])

    const editingFaq = reactive({
      id: '',
      question: '',
      answer: '',
      categoryId: ''
    })

    const editingCategory = reactive({
      id: '',
      name: ''
    })

    const configRules = {
      title: [
        { required: true, message: 'Title is required', trigger: 'blur' },
        { max: 100, message: 'Title cannot exceed 100 characters', trigger: 'blur' }
      ],
      message: [
        { required: true, message: 'Message is required', trigger: 'blur' },
        { max: 500, message: 'Message cannot exceed 500 characters', trigger: 'blur' }
      ],
      buttonText: [
        { required: true, message: 'Button text is required', trigger: 'blur' },
        { max: 50, message: 'Button text cannot exceed 50 characters', trigger: 'blur' }
      ]
    }

    const faqRules = {
      question: [
        { required: true, message: 'Question is required', trigger: 'blur' },
        { max: 200, message: 'Question cannot exceed 200 characters', trigger: 'blur' }
      ],
      answer: [
        { required: true, message: 'Answer is required', trigger: 'blur' },
        { max: 1000, message: 'Answer cannot exceed 1000 characters', trigger: 'blur' }
      ],
      categoryId: [
        { required: true, message: 'Category is required', trigger: 'change' }
      ]
    }

    const categoryRules = {
      name: [
        { required: true, message: 'Category name is required', trigger: 'blur' },
        { max: 100, message: 'Category name cannot exceed 100 characters', trigger: 'blur' }
      ]
    }

    const allCategories = computed(() => {
      return [...localCategories.value, ...newCategories.value]
    })

    const mockApiCall = async (delay = 1000) => {
      await new Promise(resolve => setTimeout(resolve, delay))
      if (Math.random() < 0.1) {
        throw new Error('API Error')
      }
      return { success: true }
    }

    const loadRemoteConfig = async () => {
      isLoadingConfig.value = true
      configError.value = null

      try {
        await mockApiCall(1200)
        const mockData = {
          EN: {
            title: 'Welcome to Our App',
            message: 'Thank you for using our application. We hope you enjoy the experience!',
            buttonText: 'Get Started',
            isEnabled: true
          },
          CN: {
            title: '欢迎使用我们的应用',
            message: '感谢您使用我们的应用程序。希望您享受这次体验！',
            buttonText: '开始使用',
            isEnabled: false
          },
          TAI: {
            title: 'ยินดีต้อนรับสู่แอปของเรา',
            message: 'ขอบคุณที่ใช้แอปพลิเคชันของเรา เราหวังว่าคุณจะสนุกกับประสบการณ์นี้!',
            buttonText: 'เริ่มต้น',
            isEnabled: true
          }
        }

        const data = mockData[remoteConfigForm.lang] || mockData.EN
        Object.assign(remoteConfigForm, data)
        lastUpdated.value = new Date()
      } catch (err) {
        configError.value = 'Failed to load remote configuration'
        console.error('Error loading remote config:', err)
      } finally {
        isLoadingConfig.value = false
      }
    }

    const handleConfigSubmit = async () => {
      if (!configFormRef.value) return

      try {
        await configFormRef.value.validate()
        isUpdatingConfig.value = true

        await mockApiCall(800)
        ElMessage.success('Remote configuration updated successfully')
        lastUpdated.value = new Date()
      } catch (err) {
        if (err !== false) {
          ElMessage.error('Failed to update remote configuration')
          console.error('Error updating config:', err)
        }
      } finally {
        isUpdatingConfig.value = false
      }
    }

    const loadFaqs = async () => {
      isLoadingFaqs.value = true
      try {
        await mockApiCall(1000)
        localFaqs.value = [
          {
            id: '1',
            question: 'How do I reset my password?',
            answer: 'You can reset your password by clicking on the "Forgot Password" link on the login page.',
            categoryId: '1'
          },
          {
            id: '2',
            question: 'How do I contact support?',
            answer: 'You can contact our support team through the help section in the app or by emailing support@example.com.',
            categoryId: '1'
          }
        ]
      } catch (err) {
        ElMessage.error('Failed to load FAQs')
        console.error('Error loading FAQs:', err)
      } finally {
        isLoadingFaqs.value = false
      }
    }

    const loadCategories = async () => {
      isLoadingCategories.value = true
      try {
        await mockApiCall(800)
        localCategories.value = [
          { id: '1', name: 'General' },
          { id: '2', name: 'Technical' },
          { id: '3', name: 'Billing' }
        ]
      } catch (err) {
        ElMessage.error('Failed to load categories')
        console.error('Error loading categories:', err)
      } finally {
        isLoadingCategories.value = false
      }
    }

    const handleAddFaq = () => {
      const newFaq = {
        question: '',
        answer: '',
        categoryId: ''
      }
      localFaqs.value.push(newFaq)
      editingFaqIndex.value = localFaqs.value.length - 1
      Object.assign(editingFaq, newFaq)
    }

    const handleEditFaq = (index) => {
      editingFaqIndex.value = index
      Object.assign(editingFaq, localFaqs.value[index])
    }

    const handleSaveFaq = async (index) => {
      if (!faqFormRef.value) return

      try {
        await faqFormRef.value.validate()
        isSavingFaq.value = true

        await mockApiCall(600)

        localFaqs.value[index] = { ...editingFaq }
        if (!localFaqs.value[index].id) {
          localFaqs.value[index].id = Date.now().toString()
        }

        editingFaqIndex.value = null
        ElMessage.success('FAQ saved successfully')
      } catch (err) {
        if (err !== false) {
          ElMessage.error('Failed to save FAQ')
          console.error('Error saving FAQ:', err)
        }
      } finally {
        isSavingFaq.value = false
      }
    }

    const handleCancelEditFaq = () => {
      if (!localFaqs.value[editingFaqIndex.value].id) {
        localFaqs.value.splice(editingFaqIndex.value, 1)
      }
      editingFaqIndex.value = null
    }

    const handleDeleteFaq = async (index) => {
      try {
        await ElMessageBox.confirm(
          'Are you sure you want to delete this FAQ?',
          'Confirm Delete',
          {
            confirmButtonText: 'Delete',
            cancelButtonText: 'Cancel',
            type: 'warning'
          }
        )

        localFaqs.value.splice(index, 1)
        ElMessage.success('FAQ deleted successfully')
      } catch {
      }
    }

    const handleAddCategory = () => {
      const tempId = `temp-${Date.now()}`
      const newCategory = {
        tempId,
        name: ''
      }
      newCategories.value.push(newCategory)
      editingCategoryId.value = tempId
      Object.assign(editingCategory, { id: tempId, name: '' })
    }

    const handleEditCategory = (category) => {
      editingCategoryId.value = category.id || category.tempId
      Object.assign(editingCategory, category)
    }

    const handleSaveCategory = async () => {
      if (!categoryFormRef.value) return

      try {
        await categoryFormRef.value.validate()
        isSavingCategory.value = true

        await mockApiCall(600)

        if (editingCategory.id.startsWith('temp-')) {
          const index = newCategories.value.findIndex(c => c.tempId === editingCategory.id)
          if (index > -1) {
            newCategories.value[index] = {
              id: Date.now().toString(),
              name: editingCategory.name
            }
          }
        } else {
          const index = localCategories.value.findIndex(c => c.id === editingCategory.id)
          if (index > -1) {
            localCategories.value[index].name = editingCategory.name
          }
        }

        editingCategoryId.value = null
        ElMessage.success('Category saved successfully')
      } catch (err) {
        if (err !== false) {
          ElMessage.error('Failed to save category')
          console.error('Error saving category:', err)
        }
      } finally {
        isSavingCategory.value = false
      }
    }

    const handleCancelEditCategory = () => {
      if (editingCategory.id.startsWith('temp-')) {
        const index = newCategories.value.findIndex(c => c.tempId === editingCategory.id)
        if (index > -1) {
          newCategories.value.splice(index, 1)
        }
      }
      editingCategoryId.value = null
    }

    const handleDeleteCategory = async (category) => {
      try {
        await ElMessageBox.confirm(
          `Are you sure you want to delete the category "${category.name}"?`,
          'Confirm Delete',
          {
            confirmButtonText: 'Delete',
            cancelButtonText: 'Cancel',
            type: 'warning'
          }
        )

        if (category.id) {
          const index = localCategories.value.findIndex(c => c.id === category.id)
          if (index > -1) {
            localCategories.value.splice(index, 1)
          }
        } else {
          const index = newCategories.value.findIndex(c => c.tempId === category.tempId)
          if (index > -1) {
            newCategories.value.splice(index, 1)
          }
        }

        ElMessage.success('Category deleted successfully')
      } catch {
      }
    }

    const getCategoryName = (categoryId) => {
      const category = [...localCategories.value, ...newCategories.value].find(c => c.id === categoryId)
      return category ? category.name : 'Unknown Category'
    }

    const formatDate = (date) => {
      return new Intl.DateTimeFormat('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      }).format(date)
    }
    onMounted(() => {
      loadRemoteConfig()
      loadFaqs()
      loadCategories()
    })

    return {

      activeTab,
      activeFaqTab,


      isLoadingConfig,
      isUpdatingConfig,
      configError,
      lastUpdated,
      configFormRef,
      remoteConfigForm,
      configRules,
      loadRemoteConfig,
      handleConfigSubmit,


      isLoadingFaqs,
      isLoadingCategories,
      isSavingFaq,
      isSavingCategory,
      editingFaqIndex,
      editingCategoryId,
      faqFormRef,
      categoryFormRef,
      localFaqs,
      localCategories,
      categories,
      allCategories,
      editingFaq,
      editingCategory,
      faqRules,
      categoryRules,
      handleAddFaq,
      handleEditFaq,
      handleSaveFaq,
      handleCancelEditFaq,
      handleDeleteFaq,
      handleAddCategory,
      handleEditCategory,
      handleSaveCategory,
      handleCancelEditCategory,
      handleDeleteCategory,
      getCategoryName,
      formatDate
    }
  }
}
</script>

<style scoped>
.page-header {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 4px 0;
}

.page-header .header-icon {
  display: flex;
  align-items: center;
  justify-content: center;
}

.page-header .header-content {
  flex: 1;
}

.page-header .header-title {
  margin: 0 0 4px 0;
  font-size: 18px;
  font-weight: 600;
  color: #303133;
  line-height: 1.4;
}

.page-header .header-description {
  margin: 0;
  font-size: 14px;
  color: #909399;
  line-height: 1.4;
}

.main-tabs {
  margin-top: 24px;

  :deep(.el-tabs__header) {
    margin: 0;
    background: transparent;
    border-bottom: 1px solid #e4e7ed;
  }

  :deep(.el-tabs__nav-wrap) {
    padding: 0;
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

.status-section {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 24px;
}

.last-updated {
  color: #909399;
  font-size: 14px;
}

.loading-container,
.error-container {
  padding: 40px 20px;
}

.config-card {
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.card-header {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.card-title {
  display: flex;
  align-items: center;
  gap: 12px;
}

.card-title h3 {
  margin: 0;
  font-size: 20px;
  font-weight: 600;
  color: #303133;
}

.config-form {
  padding: 20px;
}

.form-description {
  margin-top: 4px;
  color: #909399;
  font-size: 13px;
  line-height: 1.4;
}

.switch-container {
  background: rgba(64, 158, 255, 0.05);
  border-radius: 8px;
  padding: 16px;
}

.switch-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.switch-info h4 {
  margin: 0 0 4px 0;
  font-size: 16px;
  font-weight: 500;
  color: #303133;
}

.switch-info p {
  margin: 0;
  color: #606266;
  font-size: 14px;
}

.submit-section {
  text-align: right;
  margin-top: 24px;
}

.faqs-header {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 24px;
  text-align: center;
  justify-content: center;
}

.faqs-icon .icon {
  font-size: 48px;
  color: #409eff;
  background: rgba(64, 158, 255, 0.1);
  padding: 16px;
  border-radius: 50%;
}

.faqs-text h2 {
  margin: 0 0 8px 0;
  font-size: 32px;
  font-weight: 600;
  color: #303133;
}

.faqs-text p {
  margin: 0;
  color: #606266;
  font-size: 16px;
  max-width: 600px;
}

.faq-tabs {
  margin-top: 24px;
}

.add-section {
  text-align: right;
  margin-bottom: 24px;
}

.empty-state {
  margin: 40px 0;
}

.empty-card {
  border-radius: 12px;
}

.empty-content {
  text-align: center;
  padding: 60px 40px;
}

.empty-icon {
  font-size: 64px;
  color: #c0c4cc;
  margin-bottom: 24px;
}

.empty-content h3 {
  margin: 0 0 12px 0;
  font-size: 24px;
  font-weight: 600;
  color: #303133;
}

.empty-content p {
  margin: 0 0 24px 0;
  color: #606266;
  font-size: 16px;
  max-width: 400px;
  margin-left: auto;
  margin-right: auto;
}

.faqs-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.faq-card {
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.faq-card:hover {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
}

.faq-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.faq-actions {
  display: flex;
  gap: 8px;
}

.faq-edit,
.faq-view {
  padding: 20px;
}

.faq-question h4 {
  margin: 0 0 12px 0;
  font-size: 18px;
  font-weight: 600;
  color: #303133;
}

.faq-answer p {
  margin: 0 0 16px 0;
  color: #606266;
  line-height: 1.6;
}

.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.category-card {
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.category-card:hover {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
}

.category-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.category-actions {
  display: flex;
  gap: 8px;
}

.category-edit,
.category-view {
  padding: 20px;
}

.category-view h4 {
  margin: 0 0 8px 0;
  font-size: 18px;
  font-weight: 600;
  color: #303133;
}

.category-description {
  margin: 0;
  color: #909399;
  font-size: 14px;
}

@media (max-width: 768px) {
  .remote-config-container {
    padding: 16px;
  }

  .header-content {
    flex-direction: column;
    text-align: center;
  }

  .faqs-header {
    flex-direction: column;
  }

  .categories-grid {
    grid-template-columns: 1fr;
  }

  .switch-content {
    flex-direction: column;
    gap: 16px;
    align-items: flex-start;
  }
}
</style>
