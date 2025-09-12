<template>
  <ele-card :body-style="{ paddingBottom: '2px' }">
    <el-form label-width="100px" @keyup.enter="search" @submit.prevent="">
      <el-row :gutter="8">
        <el-col :lg="4" :md="12" :sm="12" :xs="24">
          <el-form-item label="User Name">
            <el-input clearable v-model="form.name" placeholder="Please enter user name" class="ele-fluid" />
          </el-form-item>
        </el-col>
        <el-col :lg="4" :md="12" :sm="12" :xs="24">
          <el-form-item label="Email">
            <el-input clearable v-model="form.email" placeholder="Please enter email" class="ele-fluid" />
          </el-form-item>
        </el-col>
        <el-col :lg="4" :md="12" :sm="12" :xs="24">
          <el-form-item label="Phone">
            <el-input clearable v-model="form.phoneNumber" placeholder="Please enter phone number" class="ele-fluid" />
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
import { onMounted } from 'vue'
import { useFormData } from '@/utils/use-form-data'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const props = defineProps({})
const emit = defineEmits(['search'])


const [form, resetFields] = useFormData({
  name: '',
  email: '',
  phoneNumber: ''
})


const search = () => {
  emit('search', { ...form })
}

const reset = () => {
  resetFields()
  search()
}

onMounted(() => { })
</script>
