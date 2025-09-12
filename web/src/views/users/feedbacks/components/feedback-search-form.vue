<template>
  <ele-card :body-style="{ paddingBottom: '2px' }">
    <el-form label-width="100px" @keyup.enter="search" @submit.prevent="">
      <el-row :gutter="8">
        <el-col :lg="4" :md="12" :sm="12" :xs="24">
          <el-form-item label="Search">
            <el-input clearable v-model="form.search" placeholder="Search by title, message, user name or email"
              class="ele-fluid" />
          </el-form-item>
        </el-col>
        <el-col :lg="4" :md="12" :sm="12" :xs="24">
          <el-form-item :label="t('common.enableStatusText')">
            <el-select clearable v-model="form.status" :empty-values="[null, undefined]"
              :placeholder="t('common.selectEnableStatusText')" class="ele-fluid">
              <el-option v-for="(item, index) in statusOptions" :key="index" :label="item.label" :value="item.value" />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :lg="4" :md="12" :sm="12" :xs="24">
          <el-form-item label=" Type">
            <el-select clearable v-model="form.type" placeholder="Select feedback type" class="ele-fluid">
              <el-option v-for="(item, index) in typeOptions" :key="index" :label="item.label" :value="item.value" />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :lg="4" :md="12" :sm="12" :xs="24">
          <el-form-item label-width="16px">
            <el-button type="primary" @click="search">{{ t('common.search') }}</el-button>
            <el-button @click="reset">{{ t('common.reset') }}</el-button>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </ele-card>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useFormData } from '@/utils/use-form-data'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const props = defineProps({})
const emit = defineEmits(['search'])
const [form, resetFields] = useFormData({
  search: '',
  status: '',
  type: ''
})

const statusOptions = computed(() => [
  { label: 'All', value: '' },
  { label: 'Resolved', value: 'resolved' },
  { label: 'Unresolved', value: 'unresolved' }
])

const typeOptions = computed(() => [
  { label: 'All', value: '' },
  { label: 'Bug Report', value: 'bug' },
  { label: 'Feature Request', value: 'feature' },
  { label: 'Improvement', value: 'improvement' },
  { label: 'Complaint', value: 'complaint' },
  { label: 'Compliment', value: 'compliment' },
  { label: 'Other', value: 'other' }
])

const search = () => {
  emit('search', {
    ...form
  })
}
const reset = () => {
  resetFields()
  search()
}

onMounted(() => { })
</script>
