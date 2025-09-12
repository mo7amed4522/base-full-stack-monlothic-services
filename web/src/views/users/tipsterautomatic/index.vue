<template>
  <ele-page>
    <tipster-search-form @search="onSearch" />
    <ele-card
      :flex-table="fixedHeight"
      :body-style="{ paddingBottom: '4px' }"
      :style="{
        minHeight: fixedHeight ? '380px' : void 0,
        marginBottom: fixedHeight ? '0px' : void 0
      }"
    >
      <ele-pro-table
        ref="tableRef"
        row-key="id"
        :columns="columns"
        :datasource="datasource"
        :response="response"
        v-model:current="current"
        v-model:selections="selections"
        :show-overflow-tooltip="true"
        :border="bordered"
        :toolbar="{ theme: 'default' }"
        :default-sort="{ prop: 'createdAt', order: 'descending' }"
        :footer-style="{ paddingBottom: '12px' }"
        :tableStyle="{ paddingBottom: '0', color: 'rgba(0, 0, 0, 0.65)' }"
        :loading="loading"
      >
        <template #toolbar>
          <el-button type="primary" :icon="PlusOutlined" class="ele-btn-icon" @click="showCreateDialog = true">
            {{ t('users.addTipster') }}
          </el-button>
        </template>
        <template #tipsterInfoSlot="{ row }">
          <div class="app-info-card">
            <div class="user-avatar">
              <el-avatar :size="48" :src="row.image">
                <template #default>
                  <span>{{ getInitials(row.name) }}</span>
                </template>
              </el-avatar>
            </div>
            <div class="app-details">
              <div class="app-name">{{ row.name }}</div>
              <div class="app-id">
                <span class="label">ID:</span>
                <span class="value">{{ row.id }}</span>
              </div>
              <div class="app-remark">{{ getTipsterAccuracy(row.predictions) }}% accuracy</div>
            </div>
          </div>
        </template>
        <template #predictionsSlot="{ row }">
          <div class="predictions-display">
            <span v-for="(prediction, index) in row.predictions.slice(0, 5)" :key="index" :class="['prediction-dot', prediction === 1 ? 'correct' : 'incorrect']">
              <el-icon v-if="prediction === 1">
                <Check />
              </el-icon>
              <el-icon v-else>
                <Minus />
              </el-icon>
            </span>
          </div>
        </template>
        <template #statusSlot="{ row }">
          <ele-dot :text="getTipsterAccuracy(row.predictions) >= 60 ? t('common.active') : t('common.inactive')" :type="getTipsterAccuracy(row.predictions) >= 60 ? 'success' : 'danger'" size="8px" />
        </template>
        <template #action="{ row }">
          <el-link type="primary" :underline="false" icon="edit" @click="handleEdit(row)">
            {{ t('common.edit') }}
          </el-link>
          <el-divider direction="vertical" />
          <el-link type="danger" :underline="false" icon="Delete" @click="handleDelete(row)" v-if="isDev">
            {{ t('common.deleted') }}
          </el-link>
        </template>
      </ele-pro-table>
      <div v-if="!loading && tipsters.length === 0" class="empty-state">
        <el-empty :description="t('users.noTipsterFounded')">
          <template #image>
            <el-icon size="64" color="#c0c4cc">
              <User />
            </el-icon>
          </template>
        </el-empty>
      </div>
    </ele-card>
    <el-dialog v-model="showCreateDialog" :title="t('tipsters.addTipster')" width="500px">
      <el-form :model="form" label-width="100px" ref="formRef">
        <el-form-item :label="t('tipsters.name')" prop="name" :rules="[{ required: true, message: t('tipsters.nameRequired') }]">
          <el-input v-model="form.name" :placeholder="t('tipsters.namePlaceholder')" />
        </el-form-item>
        <el-form-item :label="t('tipsters.image')">
          <div class="image-upload-section">
            <el-upload ref="uploadRef" :auto-upload="false" :show-file-list="false" accept="image/*" @change="handleImageChange">
              <div class="upload-area" :class="{ 'has-image': form.imagePreview }">
                <div v-if="form.imagePreview" class="image-preview">
                  <img :src="form.imagePreview" alt="Preview" />
                  <div class="image-overlay">
                    <el-button type="danger" size="small" @click="removeImage">
                      <el-icon>
                        <Delete />
                      </el-icon>
                    </el-button>
                  </div>
                </div>
                <div v-else class="upload-placeholder">
                  <el-icon size="48">
                    <Upload />
                  </el-icon>
                  <p>{{ t('tipsters.dragAndDropOrClick') }}</p>
                </div>
              </div>
            </el-upload>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showCreateDialog = false">{{ t('common.cancel') }}</el-button>
        <el-button type="primary" @click="handleCreate" :loading="creating">
          {{ t('tipsters.create') }}
        </el-button>
      </template>
    </el-dialog>
    <el-dialog v-model="showDeleteDialog" :title="t('tipsters.confirmDeleteTitle')" width="400px">
      <p>{{ t('tipsters.confirmDeleteDesc') }}</p>
      <template #footer>
        <el-button @click="showDeleteDialog = false">{{ t('common.cancel') }}</el-button>
        <el-button type="danger" @click="handleConfirmDelete" :loading="deleting">
          {{ t('tipsters.delete') }}
        </el-button>
      </template>
    </el-dialog>
  </ele-page>
</template>

<script setup>
  import { PlusOutlined } from '@/components/icons';
  import { EleMessage } from 'ele-admin-plus';
  import { computed, onMounted, reactive, ref } from 'vue';
  import { useI18n } from 'vue-i18n';
  import TipsterSearchForm from './components/tipster-search-form.vue';

  defineOptions({ name: 'TipsterManual' });

  const { t } = useI18n();

  const tableRef = ref(null);
  const formRef = ref(null);
  const uploadRef = ref(null);
  const bordered = ref(true);
  const fixedHeight = ref(false);
  const loading = ref(false);
  const tipsters = ref([]);
  const current = ref(null);
  const selections = ref([]);

  const lastWhere = reactive({});

  const showCreateDialog = ref(false);
  const showDeleteDialog = ref(false);
  const creating = ref(false);
  const deleting = ref(false);

  const form = reactive({
    name: '',
    image: null,
    imagePreview: null
  });

  const isDev = import.meta.env.DEV;

  const response = ref({
    dataName: 'list',
    countName: 'total'
  });

  const columns = computed(() => [
    {
      type: 'selection',
      columnKey: 'selection',
      width: 50,
      align: 'center',
      fixed: 'left'
    },
    {
      columnKey: 'tipster',
      prop: 'name',
      label: t('tipsters.manualTipster'),
      slot: 'tipsterInfoSlot',
      align: 'left',
      minWidth: 200,
      showOverflowTooltip: true
    },
    {
      columnKey: 'predictions',
      prop: 'predictions',
      label: t('tipsters.predictions'),
      slot: 'predictionsSlot',
      align: 'center',
      minWidth: 120
    },
    {
      columnKey: 'accuracy',
      prop: 'accuracy',
      label: t('tipsters.accuracy'),
      slot: 'accuracy',
      align: 'center',
      minWidth: 100
    },
    {
      columnKey: 'status',
      prop: 'status',
      label: t('tipsters.status'),
      slot: 'statusSlot',
      align: 'center',
      minWidth: 100
    },
    {
      columnKey: 'action',
      prop: 'action',
      label: t('tipsters.actions'),
      slot: 'action',
      align: 'center',
      width: 200,
      fixed: 'right'
    }
  ]);

  const reload = (where) => {
    tableRef.value?.reload?.({ page: 1, where });
  };

  const onSearch = (where) => {
    Object.assign(lastWhere, where);
    doReload();
  };

  const doReload = () => {
    reload(lastWhere);
  };

  const handleImageChange = (file) => {
    if (file.raw) {
      form.image = file.raw;
      form.imagePreview = URL.createObjectURL(file.raw);
    }
  };

  const removeImage = () => {
    form.image = null;
    form.imagePreview = null;
    if (uploadRef.value) {
      uploadRef.value.clearFiles();
    }
  };

  const handleCreate = async () => {
    if (!formRef.value) return;

    try {
      const valid = await formRef.value.validate();
      if (!valid) return;
      creating.value = true;
      EleMessage.success(t('tipsters.createdSuccess'));
      showCreateDialog.value = false;
      resetForm();
      reload();
    } catch (error) {
      EleMessage.error(t('tipsters.createFailed'));
    } finally {
      creating.value = false;
    }
  };

  const handleEdit = (row) => {
    EleMessage.info(t('tipsters.editFeatureComingSoon'));
  };

  const handleDelete = (row) => {
    current.value = row;
    showDeleteDialog.value = true;
  };

  const handleConfirmDelete = async () => {
    if (!current.value) return;

    deleting.value = true;
    try {
      EleMessage.success(t('tipsters.deletedSuccess'));
      showDeleteDialog.value = false;
      reload();
    } catch (error) {
      EleMessage.error(t('tipsters.deleteFailed'));
    } finally {
      deleting.value = false;
    }
  };

  const resetForm = () => {
    form.name = '';
    form.image = null;
    form.imagePreview = null;
    if (formRef.value) {
      formRef.value.resetFields();
    }
  };

  const generateMockTipsters = () => {
    const names = ['Manual Expert Tipster', 'Human Sports Analyst', 'Pro Manual Tipster', 'Sports Betting Expert', 'Manual Prediction Master', 'Human Analytics Pro'];

    return names.map((name, index) => {
      const predictions = Array.from({ length: 10 }, () => (Math.random() > 0.4 ? 1 : 0));
      return {
        id: (index + 4).toString(),
        name,
        image: `https://api.dicebear.com/7.x/avataaars/svg?seed=${name.replace(' ', '')}`,
        predictions,
        status: Math.random() > 0.3 ? 'active' : 'inactive',
        createdAt: new Date(Date.now() - Math.random() * 30 * 24 * 60 * 60 * 1000).toISOString(),
        type: 'manual'
      };
    });
  };

  const datasource = async ({}) => {
    loading.value = true;
    try {
      const mockTipsters = generateMockTipsters();

      let filteredTipsters = mockTipsters.filter((tipster) => tipster.type === 'manual');

      tipsters.value = filteredTipsters;

      return {
        list: filteredTipsters,
        total: filteredTipsters.length
      };
    } catch (error) {
      EleMessage.error(t('tipsters.errorLoading'));
      return { list: [], total: 0 };
    } finally {
      loading.value = false;
    }
  };

  const getTipsterAccuracy = (predictions) => {
    if (!predictions || predictions.length === 0) return 0;
    const correct = predictions.filter((p) => p === 1).length;
    return Math.round((correct / predictions.length) * 100);
  };

  onMounted(() => {
    doReload();
  });
</script>

<style lang="scss" scoped>
  .app-info-card {
    display: flex;
    align-items: center;
    gap: 12px;

    .app-icon {
      flex-shrink: 0;
    }

    .app-details {
      flex: 1;
      min-width: 0;

      .app-name {
        font-weight: 600;
        color: #303133;
        margin-bottom: 4px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }

      .app-id {
        font-size: 12px;
        color: #909399;
        margin-bottom: 2px;

        .label {
          margin-right: 4px;
        }

        .value {
          font-family: monospace;
        }
      }

      .app-remark {
        font-size: 12px;
        color: #67c23a;
        font-weight: 500;
      }
    }
  }

  .predictions-display {
    display: flex;
    gap: 4px;
    justify-content: center;

    .prediction-dot {
      width: 20px;
      height: 20px;
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 12px;

      &.correct {
        background-color: #67c23a;
        color: white;
      }

      &.incorrect {
        background-color: #f56c6c;
        color: white;
      }
    }
  }

  .image-upload-section {
    .upload-area {
      border: 2px dashed #dcdfe6;
      border-radius: 6px;
      cursor: pointer;
      position: relative;
      overflow: hidden;
      transition: border-color 0.2s;
      width: 200px;
      height: 200px;
      display: flex;
      align-items: center;
      justify-content: center;

      &:hover {
        border-color: #409eff;
      }

      &.has-image {
        border-color: #409eff;
      }

      .image-preview {
        width: 100%;
        height: 100%;
        position: relative;

        img {
          width: 100%;
          height: 100%;
          object-fit: cover;
        }

        .image-overlay {
          position: absolute;
          top: 0;
          left: 0;
          right: 0;
          bottom: 0;
          background: rgba(0, 0, 0, 0.5);
          display: flex;
          align-items: center;
          justify-content: center;
          opacity: 0;
          transition: opacity 0.2s;

          &:hover {
            opacity: 1;
          }
        }
      }

      .upload-placeholder {
        text-align: center;
        color: #8c939d;

        p {
          margin-top: 8px;
          font-size: 14px;
        }
      }
    }
  }
</style>
