<template>
  <ele-page plain hide-footer :multi-card="false">
    <ele-card style="border-radius: 0">
      <ele-text type="heading" size="lg">Advanced Form</ele-text>
      <ele-text type="placeholder" style="margin-top: 6px">Complex forms are commonly used in scenarios that require
        one-time input and submission of a large amount of data.</ele-text>
    </ele-card>
    <ele-page>
      <el-form ref="formRef" :model="form" :rules="rules" label-width="100px" @submit.prevent="">
        <ele-card header="Repository Information">
          <el-row :gutter="16">
            <el-col :lg="8" :md="12" :sm="12" :xs="24">
              <el-form-item label="Repository Name" prop="name">
                <el-input maxlength="40" clearable v-model="form.name" placeholder="Please input repository name" />
              </el-form-item>
            </el-col>
            <el-col :lg="8" :md="12" :sm="12" :xs="24">
              <el-form-item label="Repository Domain" prop="url">
                <el-input clearable v-model="form.url" placeholder="Please input repository domain">
                  <template #prepend>https://</template>
                  <template #append>.com</template>
                </el-input>
              </el-form-item>
            </el-col>
            <el-col :lg="8" :md="12" :sm="12" :xs="24">
              <el-form-item label="Repository Administrator" prop="administrator">
                <el-select clearable v-model="form.administrator" placeholder="Please select repository administrator"
                  class="ele-fluid" :popper-options="{ strategy: 'fixed' }">
                  <el-option value="1" label="SunSmile" />
                  <el-option value="2" label="Jasmine" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :lg="8" :md="12" :sm="12" :xs="24">
              <el-form-item label="Repository Approver" prop="approver">
                <el-select clearable v-model="form.approver" placeholder="Please select repository approver"
                  class="ele-fluid" :popper-options="{ strategy: 'fixed' }">
                  <el-option value="1" label="SunSmile" />
                  <el-option value="2" label="Jasmine" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :lg="8" :md="12" :sm="12" :xs="24">
              <el-form-item label="Repository Effective Time" prop="datetime">
                <el-date-picker type="datetime" v-model="form.datetime" value-format="YYYY-MM-DD HH:mm:ss"
                  placeholder="Please select repository effective time" class="ele-fluid"
                  :popper-options="{ strategy: 'fixed' }" />
              </el-form-item>
            </el-col>
            <el-col :lg="8" :md="12" :sm="12" :xs="24">
              <el-form-item label="Repository Type" prop="type">
                <el-select clearable v-model="form.type" placeholder="Please select repository type" class="ele-fluid"
                  :popper-options="{ strategy: 'fixed' }">
                  <el-option value="private" label="Private" />
                  <el-option value="public" label="Public" />
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
        </ele-card>
        <ele-card header="Task Information">
          <el-row :gutter="16">
            <el-col :lg="8" :md="12" :sm="12" :xs="24">
              <el-form-item label="Task Name" prop="task">
                <el-input clearable v-model="form.task" placeholder="Please input task name" />
              </el-form-item>
            </el-col>
            <el-col :lg="8" :md="12" :sm="12" :xs="24">
              <el-form-item label="Task Description" prop="description">
                <el-input clearable v-model="form.description" placeholder="Please input task description" />
              </el-form-item>
            </el-col>
            <el-col :lg="8" :md="12" :sm="12" :xs="24">
              <el-form-item label="Task Executor" prop="executor">
                <el-select clearable v-model="form.executor" placeholder="Please select task executor" class="ele-fluid"
                  :popper-options="{ strategy: 'fixed' }">
                  <el-option value="1" label="SunSmile" />
                  <el-option value="2" label="Jasmine" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :lg="8" :md="12" :sm="12" :xs="24">
              <el-form-item label="Task Officer" prop="officer">
                <el-select clearable v-model="form.officer" placeholder="Please select task officer" class="ele-fluid"
                  :popper-options="{ strategy: 'fixed' }">
                  <el-option value="1" label="SunSmile" />
                  <el-option value="2" label="Jasmine" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :lg="8" :md="12" :sm="12" :xs="24">
              <el-form-item label="Task Reminder" prop="reminder">
                <el-time-picker is-range range-separator="-" v-model="form.reminder" start-placeholder="Start Time"
                  end-placeholder="End Time" value-format="HH:mm:ss" class="ele-fluid"
                  :popper-options="{ strategy: 'fixed' }" />
              </el-form-item>
            </el-col>
            <el-col :lg="8" :md="12" :sm="12" :xs="24">
              <el-form-item label="Task Type" prop="taskType">
                <el-select clearable v-model="form.taskType" placeholder="Please select task type" class="ele-fluid"
                  :popper-options="{ strategy: 'fixed' }">
                  <el-option value="1" label="Private" />
                  <el-option value="2" label="Public" />
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
        </ele-card>
        <ele-card header="Task Members">
          <div style="overflow: auto">
            <ele-table style="min-width: 580px; table-layout: fixed">
              <colgroup>
                <col width="60px" />
                <col />
                <col />
                <col />
                <col width="100px" />
              </colgroup>
              <thead>
                <tr>
                  <th style="position: sticky; left: 0; z-index: 98"></th>
                  <th>Username</th>
                  <th>Employee Number</th>
                  <th>Department</th>
                  <th :style="{
                    textAlign: 'center',
                    position: 'sticky',
                    right: 0,
                    zIndex: 98
                  }">
                    Operation
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(row, index) in form.users" :key="row.key">
                  <td :style="{
                    textAlign: 'center',
                    position: 'sticky',
                    left: 0,
                    zIndex: 98
                  }">
                    {{ index + 1 }}
                  </td>
                  <td>
                    <el-form-item label-width="0px" :prop="'users.' + index + '.name'" :rules="{
                      required: true,
                      message: 'Please input username',
                      type: 'string',
                      trigger: 'blur'
                    }" class="pro-form-error-popper" style="margin-bottom: 0">
                      <el-input clearable v-model="row.name" placeholder="Please input username" />
                    </el-form-item>
                  </td>
                  <td>
                    <el-form-item label-width="0px" :prop="'users.' + index + '.number'" :rules="{
                      required: true,
                      message: 'Please input employee number',
                      type: 'string',
                      trigger: 'blur'
                    }" class="pro-form-error-popper" style="margin-bottom: 0">
                      <el-input clearable v-model="row.number" placeholder="Please input employee number" />
                    </el-form-item>
                  </td>
                  <td>
                    <el-form-item label-width="0px" :prop="'users.' + index + '.department'" :rules="{
                      required: true,
                      message: 'Please select department',
                      type: 'string',
                      trigger: 'change'
                    }" class="pro-form-error-popper" style="margin-bottom: 0">
                      <el-select clearable v-model="row.department" placeholder="Please select department"
                        class="ele-fluid" :popper-options="{ strategy: 'fixed' }">
                        <el-option value="development" label="Development" />
                        <el-option value="testing" label="Testing" />
                        <el-option value="product" label="Product" />
                      </el-select>
                    </el-form-item>
                  </td>
                  <td :style="{
                    textAlign: 'center',
                    position: 'sticky',
                    right: 0,
                    zIndex: 98
                  }">
                    <el-link type="danger" :underline="false" @click="removeUser(row, index)">Delete</el-link>
                  </td>
                </tr>
                <tr v-if="!form.users || !form.users.length">
                  <td colspan="5" style="text-align: center">
                    <ele-text style="padding: 4px 0" type="secondary">No data available</ele-text>
                  </td>
                </tr>
              </tbody>
            </ele-table>
          </div>
          <el-button :icon="PlusOutlined" style="margin-top: 16px; width: 100%" @click="addUser">
            <span style="padding-top: 1px">Add member</span>
          </el-button>
        </ele-card>
      </el-form>
    </ele-page>
    <!-- 底部工具栏 -->
    <ele-bottom-bar>
      <ele-text v-if="validMsg" type="danger" :icon="CloseCircleOutlined">
        <span>{{ validMsg }}</span>
      </ele-text>
      <template #extra>
        <el-button type="primary" :loading="loading" @click="submit">Submit</el-button>
      </template>
    </ele-bottom-bar>
  </ele-page>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { CloseCircleOutlined, PlusOutlined } from '@/components/icons';
import { useFormData } from '@/utils/use-form-data';
import { listAddedUsers } from '@/api/example';

/** 加载状态 */
const loading = ref(false);

/** 表单实例 */
const formRef = ref(null);

/** 表单数据 */
const [form, resetFields] = useFormData({
  name: '',
  url: '',
  administrator: void 0,
  approver: void 0,
  datetime: '',
  type: void 0,
  task: '',
  description: '',
  executor: void 0,
  officer: void 0,
  reminder: void 0,
  taskType: void 0,
  users: []
});

/** 表单验证规则 */
const rules = reactive({
  name: [
    {
      required: true,
      message: 'Please input repository name',
      type: 'string',
      trigger: 'blur'
    }
  ],
  url: [
    {
      required: true,
      message: 'Please input repository domain',
      type: 'string',
      trigger: 'blur'
    }
  ],
  administrator: [
    {
      required: true,
      message: 'Please select repository administrator',
      type: 'string',
      trigger: 'blur'
    }
  ],
  approver: [
    {
      required: true,
      message: 'Please select approver',
      type: 'string',
      trigger: 'blur'
    }
  ],
  datetime: [
    {
      required: true,
      message: 'Please select effective time',
      type: 'string',
      trigger: 'blur'
    }
  ],
  type: [
    {
      required: true,
      message: 'Please select repository type',
      type: 'string',
      trigger: 'blur'
    }
  ],
  task: [
    {
      required: true,
      message: 'Please input task name',
      type: 'string',
      trigger: 'blur'
    }
  ],
  description: [
    {
      required: true,
      message: 'Please input task description',
      type: 'string',
      trigger: 'blur'
    }
  ],
  executor: [
    {
      required: true,
      message: 'Please select executor',
      type: 'string',
      trigger: 'blur'
    }
  ],
  officer: [
    {
      required: true,
      message: 'Please select officer',
      type: 'string',
      trigger: 'blur'
    }
  ],
  reminder: [
    {
      required: true,
      message: 'Please select reminder time',
      type: 'array',
      trigger: 'change'
    }
  ],
  taskType: [
    {
      required: true,
      message: 'Please select task type',
      type: 'string',
      trigger: 'blur'
    }
  ]
});

/** 表单验证失败提示信息 */
const validMsg = ref('');

/** 表单提交 */
const submit = () => {
  formRef.value?.validate?.((valid, obj) => {
    if (!valid) {
      const errors = obj ? Object.keys(obj).length : 0;
      validMsg.value = `Total ${errors} items failed validation`;
      return;
    }
    validMsg.value = '';
    loading.value = true;
    setTimeout(() => {
      loading.value = false;
      EleMessage.success('Submit success');
      resetFields();
    }, 1000);
  });
};

/** 添加 */
const addUser = () => {
  form.users.push({
    key: Date.now() + '-' + form.users.length,
    number: '',
    name: '',
    department: void 0
  });
};

/** 删除 */
const removeUser = (_row, index) => {
  form.users.splice(index, 1);
};

/** 查询已添加 */
listAddedUsers()
  .then((data) => {
    form.users = data;
  })
  .catch((e) => {
    EleMessage.error(e.message);
  });
</script>

<script>
export default {
  name: 'FormAdvanced'
};
</script>
