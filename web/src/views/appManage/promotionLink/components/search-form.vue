<!-- Search form -->
<template>
  <ele-card :body-style="{ paddingBottom: '2px' }">
    <el-form label-width="100px" @keyup.enter="search" @submit.prevent="">
      <el-row :gutter="8">
        <el-col :lg="4" :md="12" :sm="12" :xs="24">
          <el-form-item label="Application">
            <el-select clearable v-model="form.appId" :empty-values="[null, undefined]" placeholder="Select application"
              class="ele-fluid">
              <el-option v-for="item in appList" :key="item.id" :label="item.appName" :value="item.id" />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :lg="4" :md="12" :sm="12" :xs="24">
          <el-form-item label="Link name">
            <el-input clearable v-model="form.name" placeholder="Input link name" class="ele-fluid" />
          </el-form-item>
        </el-col>
        <el-col :lg="4" :md="12" :sm="12" :xs="24">
          <el-form-item label="Domain">
            <el-select clearable v-model="form.domainId" :empty-values="[null, undefined]" placeholder="Select domain"
              class="ele-fluid">
              <el-option v-for="item in domainList" :key="item.id" :label="item.domain" :value="item.id" />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :lg="4" :md="12" :sm="12" :xs="24">
          <el-form-item label="Template">
            <el-select clearable v-model="form.templateId" :empty-values="[null, undefined]"
              placeholder="Select template" class="ele-fluid">
              <el-option v-for="item in templateList" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :lg="4" :md="12" :sm="12" :xs="24">
          <el-form-item label="Is allow edit URL" label-width="120px">
            <el-select clearable v-model="form.isAllowEditUrl" :empty-values="[null, undefined]" placeholder="Select"
              class="ele-fluid">
              <el-option v-for="item in isAllowEditUrlOptions" :key="item.value" :label="item.label"
                :value="item.value" />
            </el-select>
          </el-form-item>
        </el-col>
        <!-- <el-col :lg="6" :md="12" :sm="12" :xs="24">
          <el-form-item :label="t('common.createTimeText')">
            <el-date-picker
              unlink-panels
              type="datetimerange"
              v-model="dateRange"
              range-separator="-"
              format="YYYY-MM-DD HH:mm"
              value-format="YYYY-MM-DD HH:mm"
              :start-placeholder="t('common.beginTimeText')"
              :end-placeholder="t('common.endTimeText')"
              class="ele-fluid"
            />
          </el-form-item>
        </el-col> -->

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
import { ref, onMounted, computed } from 'vue';
import { useFormData } from '@/utils/use-form-data';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

const props = defineProps({
  appList: {
    type: Array,
    default: () => []
  },
  domainList: {
    type: Array,
    default: () => []
  },
  templateList: {
    type: Array,
    default: () => []
  }
});
const emit = defineEmits(['search']);

/** Search form data */
const [form, resetFields] = useFormData({
  name: '',
  appId: null,
  domainId: null,
  templateId: null,
  isAllowEditUrl: null, // Is it allowed to modify the URL? 1 Allow -1 Do not allow
});

const isAllowEditUrlOptions = computed(() => [
  { label: 'Allow', value: 1 },
  { label: 'Do not allow', value: -1 }
]);

/** Date range */
const dateRange = ref(['', '']);

/** Search */
const search = () => {
  const [d1, d2] = dateRange.value ?? [];
  emit('search', {
    ...form,
    startTime: d1 ? `${d1}` : '',
    endTime: d2 ? `${d2}` : ''
  });
};

/**  Reset */
const reset = () => {
  resetFields();
  dateRange.value = ['', ''];
  search();
};

onMounted(() => { });
</script>
