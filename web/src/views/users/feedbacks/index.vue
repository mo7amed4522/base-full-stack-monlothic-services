<template>
  <ele-page>
    <feedback-search-form @search="onSearch" />
    <ele-card :body-style="{ paddingBottom: '4px' }">
      <ele-pro-table ref="tableRef" row-key="id" :columns="columns" :datasource="datasource" :response="response"
        v-model:selections="selections" :show-overflow-tooltip="true" :border="true" :toolbar="{ theme: 'default' }"
        :default-sort="{ prop: 'createdAt', order: 'descending' }" :footer-style="{ paddingBottom: '12px' }"
        :tableStyle="{ paddingBottom: '0', color: 'rgba(0, 0, 0, 0.65)' }" :loading="loading">
        <template #toolbar>
          <div class="feedback-stats">
            <el-tag type="info" size="large" style="margin-right: 12px;">
              {{ t('feedbacks.totalFeedbacks', { count: totalFeedbacks }) }}
            </el-tag>
            <el-tag type="success" size="large" style="margin-right: 12px;">
              {{ t('feedbacks.resolved', { count: resolvedCount }) }}
            </el-tag>
            <el-tag type="warning" size="large">
              {{ t('feedbacks.unresolved', { count: unresolvedCount }) }}
            </el-tag>
          </div>
        </template>
        <template #feedbackInfoSlot="{ row }">
          <div class="user-cell">
            <el-avatar :size="36" :src="row.user?.image">
              {{ row.user?.name?.charAt(0) || 'U' }}
            </el-avatar>
            <div class="user-info">
              <div class="user-name">{{ row.user?.name || 'Anonymous' }}</div>
              <div class="user-contact">{{ row.user?.email || 'N/A' }}</div>
            </div>
          </div>
        </template>
        <template #feedbackTypeSlot="{ row }">
          <el-tag :type="getFeedbackTypeColor(row.type)">
            {{ getFeedbackTypeLabel(row.type) }}
          </el-tag>
        </template>
        <template #feedbackContentSlot="{ row }">
          <div class="feedback-content">
            <div class="feedback-title">{{ row.title }}</div>
            <div class="feedback-message">{{ truncateText(row.message, 100) }}</div>
            <div class="feedback-meta">
              <el-tag size="small" v-if="row.rating">
                <el-icon>
                  <Star />
                </el-icon>
                {{ row.rating }}/5
              </el-tag>
            </div>
          </div>
        </template>
        <template #statusSlot="{ row }">
          <el-tag :type="row.status === 'resolved' ? 'success' : 'warning'" size="large">
            {{ row.status === 'resolved' ? t('feedbacks.resolved') : t('feedbacks.unresolved') }}
          </el-tag>
        </template>
        <template #action="{ row }">
          <el-link type="primary" :underline="false" @click="handleViewFeedback(row)">
            <el-icon>
              <View />
            </el-icon>
            {{ t('common.view') }}
          </el-link>
          <el-divider direction="vertical" />
          <el-link :type="row.status === 'resolved' ? 'warning' : 'success'" :underline="false"
            @click="handleToggleStatus(row)">
            <el-icon>
              <Check />
            </el-icon>
            {{ row.status === 'resolved' ? t('feedbacks.markUnresolved') : t('feedbacks.markResolved') }}
          </el-link>
          <el-divider direction="vertical" v-if="isDev" />
          <el-link type="danger" :underline="false" @click="handleDeleteFeedback(row)" v-if="isDev">
            <el-icon>
              <Delete />
            </el-icon>
            {{ t('common.delete') }}
          </el-link>
        </template>
      </ele-pro-table>
      <div v-if="!loading && feedbacks.length === 0" class="empty-state">
        <el-empty :description="t('feedbacks.noFeedbacksFound')">
          <template #image>
            <el-icon size="64" color="#c0c4cc">
              <MessageBox />
            </el-icon>
          </template>
        </el-empty>
      </div>
    </ele-card>
    <el-dialog v-model="showFeedbackDialog" :title="t('feedbacks.viewFeedback')" width="600px"
      :close-on-click-modal="false">
      <div v-if="currentFeedback" class="feedback-dialog-content">
        <div class="feedback-header">
          <div class="user-info">
            <el-avatar :size="40" :src="currentFeedback.user?.image">
              <template #default>
                <span>{{ getInitials(currentFeedback.user?.name) }}</span>
              </template>
            </el-avatar>
            <div class="user-details">
              <div class="user-name">{{ currentFeedback.user?.name || 'Anonymous' }}</div>
              <div class="user-email">{{ currentFeedback.user?.email || 'N/A' }}</div>
            </div>
          </div>
          <div class="feedback-meta">
            <el-tag :type="getFeedbackTypeColor(currentFeedback.type)">
              {{ getFeedbackTypeLabel(currentFeedback.type) }}
            </el-tag>
            <el-tag v-if="currentFeedback.rating">
              <el-icon>
                <Star />
              </el-icon>
              {{ currentFeedback.rating }}/5
            </el-tag>
          </div>
        </div>
        <el-divider />
        <div class="feedback-content-full">
          <h4>{{ currentFeedback.title }}</h4>
          <p>{{ currentFeedback.message }}</p>
        </div>
        <el-divider />
        <div class="feedback-footer">
          <div class="feedback-date">
            <el-icon>
              <Calendar />
            </el-icon>
            {{ formatDate(currentFeedback.createdAt) }}
          </div>
          <div class="feedback-status">
            <ele-dot :text="currentFeedback.status === 'resolved' ? t('feedbacks.resolved') : t('feedbacks.unresolved')"
              :type="currentFeedback.status === 'resolved' ? 'success' : 'warning'" size="8px" />
          </div>
        </div>
      </div>
      <template #footer>
        <el-button @click="showFeedbackDialog = false">{{ t('common.close') }}</el-button>
        <el-button :type="currentFeedback?.status === 'resolved' ? 'warning' : 'success'"
          @click="handleToggleStatusFromDialog" :loading="updatingStatus">
          {{ currentFeedback?.status === 'resolved' ? t('feedbacks.markUnresolved') : t('feedbacks.markResolved') }}
        </el-button>
      </template>
    </el-dialog>
    <el-dialog v-model="showDeleteDialog" :title="t('feedbacks.confirmDeleteTitle')" width="400px">
      <p>{{ t('feedbacks.confirmDeleteDesc') }}</p>
      <template #footer>
        <el-button @click="showDeleteDialog = false">{{ t('common.cancel') }}</el-button>
        <el-button type="danger" @click="handleConfirmDelete" :loading="deletingFeedback">
          {{ t('feedbacks.delete') }}
        </el-button>
      </template>
    </el-dialog>
  </ele-page>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { EleMessage } from 'ele-admin-plus'
import { View, Check, Delete, Star, Calendar, MessageBox } from '@element-plus/icons-vue'
import FeedbackSearchForm from './components/feedback-search-form.vue'

defineOptions({ name: 'UsersFeedbacks' })

const { t } = useI18n()
const tableRef = ref(null)
const loading = ref(false)
const feedbacks = ref([])
const selections = ref([])

const showFeedbackDialog = ref(false)
const showDeleteDialog = ref(false)
const currentFeedback = ref(null)
const updatingStatus = ref(false)
const deletingFeedback = ref(false)

const lastWhere = reactive({})

const isDev = import.meta.env.DEV

const response = ref({
  dataName: 'list',
  countName: 'total'
})

const totalFeedbacks = computed(() => feedbacks.value.length)
const resolvedCount = computed(() => feedbacks.value.filter(f => f.status === 'resolved').length)
const unresolvedCount = computed(() => feedbacks.value.filter(f => f.status === 'unresolved').length)

const getInitials = (name) => {
  if (!name) return 'U'
  return name.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2)
}

const truncateText = (text, maxLength) => {
  if (!text) return ''
  return text.length > maxLength ? text.substring(0, maxLength) + '...' : text
}

const getFeedbackTypeColor = (type) => {
  const colors = {
    'bug': 'danger',
    'feature': 'primary',
    'improvement': 'success',
    'complaint': 'warning',
    'compliment': 'success',
    'other': 'info'
  }
  return colors[type] || 'info'
}

const getFeedbackTypeLabel = (type) => {
  return t(`feedbacks.types.${type}`, type)
}

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleDateString()
}

const reload = (where) => {
  tableRef.value?.reload?.({ page: 1, where })
}

const onSearch = (where) => {
  Object.assign(lastWhere, where)
  reload(lastWhere)
}

const doReload = () => {
  reload(lastWhere)
}

const handleViewFeedback = (feedback) => {
  currentFeedback.value = feedback
  showFeedbackDialog.value = true
}

const handleToggleStatus = async (feedback) => {
  updatingStatus.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    const newStatus = feedback.status === 'resolved' ? 'unresolved' : 'resolved'
    feedback.status = newStatus
    EleMessage.success(t('feedbacks.statusUpdated'))
    doReload()
  } catch (error) {
    EleMessage.error(t('feedbacks.statusUpdateFailed'))
  } finally {
    updatingStatus.value = false
  }
}

const handleToggleStatusFromDialog = async () => {
  if (!currentFeedback.value) return

  await handleToggleStatus(currentFeedback.value)
  showFeedbackDialog.value = false
}

const handleDeleteFeedback = (feedback) => {
  currentFeedback.value = feedback
  showDeleteDialog.value = true
}

const handleConfirmDelete = async () => {
  if (!currentFeedback.value) return

  deletingFeedback.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))

    EleMessage.success(t('feedbacks.deletedSuccess'))
    showDeleteDialog.value = false
    doReload()
  } catch (error) {
    EleMessage.error(t('feedbacks.deleteFailed'))
  } finally {
    deletingFeedback.value = false
  }
}

const generateFakeUsers = () => {
  const names = [
    'Khaled Mohamed', 'Ahmed Mohamed', 'Khaled Mohamed', 'Ahmed Mohamed',
    'ALi Ali', 'Ali Ali', 'Khaled Mohamed', 'Khaled Mohamed',
    'Ahmed Mohamed', 'Ahmed Mohamed', 'Khaled Ahmed', 'Ahmed Mohamed'
  ]
  const emails = [
    'khaled@example.com', 'ahmed@example.com', 'mohamed@example.com', 'khaled@example.com',
    'mo7amed4522@example.com', 'khaled@example.com', 'ali@example.com', 'ahmed@example.com',
    'ahmed4522@example.com', 'tom@example.com', 'mohamed@example.com', 'mohamed@example.com'
  ]

  return names.map((name, index) => ({
    id: index + 1,
    name,
    email: emails[index],
    image: `https://api.dicebear.com/7.x/avataaars/svg?seed=${name.replace(' ', '')}`
  }))
}

const columns = computed(() => [
  {
    type: 'selection',
    width: 45,
    align: 'center',
    fixed: 'left'
  },
  {
    columnKey: 'feedbackInfo',
    prop: 'user.name',
    label: t('feedbacks.user'),
    slot: 'feedbackInfoSlot',
    align: 'left',
    minWidth: 200
  },
  {
    columnKey: 'feedbackType',
    prop: 'type',
    label: t('feedbacks.type'),
    slot: 'feedbackTypeSlot',
    align: 'center',
    minWidth: 120
  },
  {
    columnKey: 'feedbackContent',
    prop: 'title',
    label: t('feedbacks.feedback'),
    slot: 'feedbackContentSlot',
    align: 'left',
    minWidth: 300
  },
  {
    columnKey: 'status',
    prop: 'status',
    label: t('feedbacks.status'),
    slot: 'statusSlot',
    align: 'center',
    minWidth: 100
  },
  {
    columnKey: 'createdAt',
    prop: 'createdAt',
    label: t('feedbacks.date'),
    align: 'center',
    minWidth: 150,
    formatter: (row) => formatDate(row.createdAt)
  },
  {
    columnKey: 'action',
    prop: 'action',
    label: t('feedbacks.actions'),
    slot: 'action',
    align: 'center',
    width: 250,
    fixed: 'right'
  }
])


const datasource = async ({ where }) => {
  loading.value = true
  try {
    const users = generateFakeUsers()
    const mockFeedbacks = [
      {
        id: '1',
        title: 'App crashes on startup',
        message: 'The application crashes immediately when I try to open it on my iPhone 12. This has been happening for the past week.',
        type: 'bug',
        status: 'unresolved',
        rating: 2,
        user: users[0],
        createdAt: '2024-01-15T10:30:00Z'
      },
      {
        id: '2',
        title: 'Great new feature request',
        message: 'It would be amazing if you could add dark mode support. Many users have been requesting this feature.',
        type: 'feature',
        status: 'resolved',
        rating: 5,
        user: users[1],
        createdAt: '2024-01-14T14:20:00Z'
      },
      {
        id: '3',
        title: 'Slow loading times',
        message: 'The app takes too long to load content, especially on slower internet connections.',
        type: 'improvement',
        status: 'unresolved',
        rating: 3,
        user: users[2],
        createdAt: '2024-01-13T09:15:00Z'
      },
      {
        id: '4',
        title: 'Love the new update!',
        message: 'The latest update is fantastic! The new UI is much more intuitive and the performance improvements are noticeable.',
        type: 'compliment',
        status: 'resolved',
        rating: 5,
        user: users[3],
        createdAt: '2024-01-12T16:45:00Z'
      },
      {
        id: '5',
        title: 'Payment processing issue',
        message: 'I am unable to complete my purchase. The payment keeps failing even though my card is valid.',
        type: 'complaint',
        status: 'unresolved',
        rating: 1,
        user: users[4],
        createdAt: '2024-01-11T11:30:00Z'
      }
    ]

    let filteredFeedbacks = mockFeedbacks

    if (where?.search) {
      const searchTerm = where.search.toLowerCase()
      filteredFeedbacks = filteredFeedbacks.filter(feedback =>
        feedback.title.toLowerCase().includes(searchTerm) ||
        feedback.message.toLowerCase().includes(searchTerm) ||
        feedback.user?.name?.toLowerCase().includes(searchTerm) ||
        feedback.user?.email?.toLowerCase().includes(searchTerm)
      )
    }

    if (where?.status && where.status !== 'all') {
      filteredFeedbacks = filteredFeedbacks.filter(feedback => feedback.status === where.status)
    }

    if (where?.type && where.type !== 'all') {
      filteredFeedbacks = filteredFeedbacks.filter(feedback => feedback.type === where.type)
    }

    if (where?.rating) {
      filteredFeedbacks = filteredFeedbacks.filter(feedback => feedback.rating >= where.rating)
    }

    if (where?.sort === 'createdAt:asc') {
      filteredFeedbacks.sort((a, b) => new Date(a.createdAt) - new Date(b.createdAt))
    } else {
      filteredFeedbacks.sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt))
    }

    feedbacks.value = filteredFeedbacks

    return {
      list: filteredFeedbacks,
      total: filteredFeedbacks.length
    }
  } catch (error) {
    EleMessage.error(t('feedbacks.errorLoading'))
    return {
      list: [],
      total: 0
    }
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  doReload()
})
</script>

<style scoped>
.user-cell {
  display: flex;
  align-items: center;
  gap: 12px;
}

.user-info {
  flex: 1;
}

.user-name {
  font-weight: 500;
  color: #333;
}

.user-contact {
  font-size: 12px;
  color: #666;
  margin-top: 2px;
}

.feedback-content {
  max-width: 300px;
}

.feedback-title {
  font-weight: 500;
  color: #303133;
  margin-bottom: 4px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.feedback-message {
  font-size: 12px;
  color: #606266;
  line-height: 1.4;
  margin-bottom: 4px;
}

.feedback-meta {
  display: flex;
  align-items: center;
  gap: 8px;
}

.empty-state {
  text-align: center;
  padding: 60px 20px;
  color: #666;
}

.feedback-dialog-content {
  padding: 0;
}

.feedback-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 16px;
}

.user-details {
  flex: 1;
}

.user-email {
  font-size: 12px;
  color: #909399;
}

.feedback-content-full h4 {
  margin: 0 0 12px 0;
  color: #303133;
  font-size: 16px;
}

.feedback-content-full p {
  margin: 0;
  color: #606266;
  line-height: 1.6;
}

.feedback-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.feedback-date {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  color: #909399;
}

.feedback-status {
  display: flex;
  align-items: center;
}
</style>
