<template>
  <ele-page>
    <ele-card>
      <el-row>
        <el-col :sm="8" :xs="24">
          <div style="text-align: center">
            <ele-text type="placeholder">In progress tasks</ele-text>
            <ele-text size="xl" style="margin-top: 6px">10 tasks</ele-text>
          </div>
        </el-col>
        <el-col :sm="8" :xs="24">
          <div style="text-align: center">
            <ele-text type="placeholder">Remaining tasks</ele-text>
            <ele-text size="xl" style="margin-top: 6px">3 tasks</ele-text>
          </div>
        </el-col>
        <el-col :sm="8" :xs="24">
          <div style="text-align: center">
            <ele-text type="placeholder">Total task duration</ele-text>
            <ele-text size="xl" style="margin-top: 6px">120 hours</ele-text>
          </div>
        </el-col>
      </el-row>
    </ele-card>
    <ele-card>
      <!-- 头部工具栏 -->
      <ele-toolbar title="Advanced list" :title-props="{ size: 'md' }">
        <template #tools>
          <div style="display: flex; align-items: center">
            <el-radio-group v-model="where.status" @change="query">
              <el-radio-button value="0" label="All" />
              <el-radio-button value="1" label="In progress" />
              <el-radio-button value="2" label="Completed" />
            </el-radio-group>
            <div style="width: 200px; margin-left: 12px" class="hidden-xs-only">
              <el-input v-model="where.keyword" placeholder="Search">
                <template #append>
                  <el-button :icon="SearchOutlined" class="ele-btn-icon" @click="query" />
                </template>
              </el-input>
            </div>
          </div>
        </template>
      </ele-toolbar>
      <div style="margin-top: 8px">
        <el-button :icon="PlusOutlined" class="ele-fluid" @click="openEdit()">
          <span style="transform: translateY(1px)">Add</span>
        </el-button>
      </div>
      <!-- 数据列表 -->
      <ele-loading :loading="loading" style="min-height: 100px">
        <div v-for="item in data" :key="item.id">
          <div class="list-item">
            <div class="list-item-avatar">
              <el-avatar :size="60" shape="square" :src="item.cover" />
              <div class="list-item-avatar-extra">
                <div style="margin-bottom: 2px">{{ item.title }}</div>
                <ele-text size="sm" type="placeholder">
                  {{ item.content }}
                </ele-text>
              </div>
            </div>
            <div class="list-item-owner">
              <div class="list-item-title">Creator</div>
              <ele-text size="sm" type="placeholder">{{ item.user }}</ele-text>
            </div>
            <div class="list-item-time">
              <div class="list-item-title">Start time</div>
              <ele-text size="sm" type="placeholder">{{ item.startTime }}</ele-text>
            </div>
            <div class="list-item-progress" style="display: block">
              <el-progress :status="item.status" :percentage="item.progress" />
            </div>
            <div class="list-item-tools">
              <el-link type="primary" :underline="false" @click="openEdit(item)">Edit</el-link>
              <el-divider direction="vertical" />
              <ele-dropdown :items="[
                { title: 'Share', command: 'share' },
                { title: 'Delete', command: 'remove', danger: true }
              ]" @command="(key) => dropClick(key, item)">
                <el-link type="primary" :underline="false">
                  <span>More</span>
                  <el-icon :size="12">
                    <ArrowDown />
                  </el-icon>
                </el-link>
              </ele-dropdown>
            </div>
          </div>
          <el-divider style="margin: 0; opacity: 0.6" />
        </div>
      </ele-loading>
      <ele-pagination :total="count" v-model:page-size="limit" v-model:current-page="page"
        layout="prev, pager, next, sizes, jumper" :pageSizes="[5, 10, 20]"
        style="margin-top: 18px; justify-content: center" />
    </ele-card>
    <!-- 编辑弹窗 -->
    <ele-modal form :width="460" v-model="visible" :title="form.id ? 'Edit task' : 'Add task'">
      <el-form ref="formRef" :model="form" :rules="rules" label-width="88px" @submit.prevent="">
        <el-form-item label="Task name" prop="title">
          <el-input clearable :maxlength="20" v-model="form.title" placeholder="Please input task name" />
        </el-form-item>
        <el-form-item label="Start time" prop="time">
          <el-date-picker type="datetime" v-model="form.time" format="YYYY-MM-DD HH:mm" value-format="YYYY-MM-DD HH:mm"
            placeholder="Please select start time" class="ele-fluid" />
        </el-form-item>
        <el-form-item label="Assignee" prop="user">
          <el-select clearable v-model="form.user" placeholder="Please select assignee" class="ele-fluid">
            <el-option value="SunSmile" label="SunSmile" />
            <el-option value="Pojin" label="Pojin" />
            <el-option value="SuperWill" label="SuperWill" />
            <el-option value="Jasmine" label="Jasmine" />
            <el-option value="Vast" label="Vast" />
          </el-select>
        </el-form-item>
        <el-form-item label="Task description">
          <el-input :rows="4" type="textarea" v-model="form.content" placeholder="Please input task description" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="closeEdit">Cancel</el-button>
        <el-button type="primary" :loading="submitLoading" @click="onSubmit">Save</el-button>
      </template>
    </ele-modal>
  </ele-page>
</template>

<script setup>
import { ref, reactive, watch } from 'vue';
import { ElMessageBox } from 'element-plus/es';
import { EleMessage } from 'ele-admin-plus/es';
import { SearchOutlined, PlusOutlined, ArrowDown } from '@/components/icons';
import { useFormData } from '@/utils/use-form-data';

/** 列表加载状态 */
const loading = ref(false);

/** 列表数据 */
const data = ref([]);

/** 搜索表单 */
const where = reactive({
  status: '0',
  keyword: ''
});

/** 第几页 */
const page = ref(1);

/** 每页多少条 */
const limit = ref(5);

/** 总数量 */
const count = ref(0);

/** 编辑弹窗是否显示 */
const visible = ref(false);

/** 表单实例 */
const formRef = ref(null);

/** 编辑弹窗表单数据 */
const [form, resetFields, assignFields] = useFormData({
  id: void 0,
  title: '',
  time: void 0,
  user: '',
  content: ''
});

/** 编辑弹窗表单验证规则 */
const rules = reactive({
  title: [
    {
      required: true,
      message: 'Please input task name',
      type: 'string',
      trigger: 'blur'
    }
  ],
  time: [
    {
      required: true,
      message: 'Please select start time',
      type: 'string',
      trigger: 'blur'
    }
  ],
  user: [
    {
      required: true,
      message: 'Please select assignee',
      type: 'string',
      trigger: 'blur'
    }
  ]
});

/** 编辑表单提交状态 */
const submitLoading = ref(false);

/** 查询数据 */
const query = () => {
  loading.value = true;
  setTimeout(() => {
    loading.value = false;
    count.value = 20;
    data.value = [
      {
        id: 1,
        title: 'ElementUI',
        time: '2020-06-13 08:33:12',
        user: 'SunSmile',
        progress: 87,
        content: 'Element, a Vue 2.0 component library for developers, designers, and product managers, provides supporting design resources to help you quickly build your website.',
        cover: 'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg'
      },
      {
        id: 2,
        title: 'Vue.js',
        time: '2020-06-13 06:40:13',
        user: 'Pojin',
        progress: 100,
        status: 'success',
        content: 'Vue is a progressive framework for building user interfaces. Unlike other monolithic frameworks, Vue is designed to be incrementally adoptable. The core library is focused on the view layer only, and is easy to pick up and integrate with other libraries or existing projects.',
        cover: 'https://cdn.eleadmin.com/20200609/b6a811873e704db49db994053a5019b2.jpg'
      },
      {
        id: 3,
        title: 'Vuex',
        time: '2020-06-13 04:40:20',
        user: 'SuperWill',
        progress: 75,
        content: 'Vuex is a state management pattern + library for Vue.js applications. It serves as a centralized store for all the components in an application, with rules ensuring that the state can only be mutated in a predictable fashion.',
        cover: 'https://cdn.eleadmin.com/20200609/948344a2a77c47a7a7b332fe12ff749a.jpg'
      },
      {
        id: 4,
        title: 'Vue Router',
        time: '2020-06-13 02:40:05',
        user: 'Jasmine',
        progress: 65,
        content: 'Vue Router is the official router for Vue.js. It deeply integrates with Vue.js core to make building single-page applications with Vue.js a breeze.',
        cover: 'https://cdn.eleadmin.com/20200609/f6bc05af944a4f738b54128717952107.jpg'
      },
      {
        id: 5,
        title: 'Sass',
        time: '2020-06-13 00:40:58',
        user: 'Vast',
        progress: 45,
        status: 'exception',
        content: 'Sass is a preprocessor scripting language that is interpreted or compiled into Cascading Style Sheets (CSS).',
        cover: 'https://cdn.eleadmin.com/20200609/2d98970a51b34b6b859339c96b240dcd.jpg'
      }
    ];
  }, 300);
};

/** 显示编辑弹窗 */
const openEdit = (row) => {
  resetFields();
  formRef.value?.clearValidate?.();
  visible.value = true;
  if (row) {
    assignFields(row);
  }
};

/** 关闭编辑弹窗 */
const closeEdit = () => {
  visible.value = false;
};

/** 保存编辑 */
const onSubmit = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    submitLoading.value = true;
    setTimeout(() => {
      closeEdit();
      submitLoading.value = false;
      EleMessage.success('Save successful');
      if (form.id) {
        // 保存修改
        const index = data.value.findIndex((d) => d.id === form.id);
        if (index !== -1) {
          Object.assign(data.value[index], form);
        }
      } else {
        // 保存添加
        data.value.push({
          ...form,
          id: Date.now(),
          cover: 'https://cdn.eleadmin.com/20200610/RZ8FQmZfHkcffMlTBCJllBFjEhEsObVo.jpg'
        });
      }
    }, 300);
  });
};

/** 下拉菜单点击事件 */
const dropClick = (key, item) => {
  if (key === 'remove') {
    // 删除
    ElMessageBox.confirm(`Are you sure you want to delete "${item.title}"?`, 'System prompts', {
      type: 'warning',
      draggable: true
    })
      .then(() => {
        EleMessage.success('Delete successful');
      })
      .catch(() => { });
  } else if (key === 'share') {
    EleMessage.success('Share successful');
  }
};

// query();

watch([page, limit], () => {
  query();
});
</script>

<script>
export default {
  name: 'ListAdvanced'
};
</script>

<style lang="scss" scoped>
/* 列表样式 */
.list-item {
  display: flex;
  align-items: center;
  padding: 16px 8px;

  .list-item-avatar {
    flex: 1;
    display: flex;
    align-items: center;

    :deep(.el-avatar) {
      flex-shrink: 0;
    }

    .list-item-avatar-extra {
      flex: 1;
      padding-left: 12px;
      box-sizing: border-box;
    }
  }

  &>div+div {
    margin-left: 20px;
    flex-shrink: 0;
  }

  .list-item-owner {
    width: 80px;
  }

  .list-item-time {
    width: 160px;
  }

  .list-item-progress {
    width: 180px;
  }

  .list-item-title {
    margin-bottom: 4px;
  }

  .list-item-tools {
    display: flex;
    align-items: center;
  }
}

/* 响应式 */
@media screen and (max-width: 1340px) {
  .list-item {
    &>div+div {
      margin-left: 10px;
    }

    .list-item-owner {
      width: 70px;
    }

    .list-item-time {
      width: 140px;
    }

    .list-item-progress {
      width: 100px;
    }
  }
}

@media screen and (max-width: 1100px) {
  .list-item {
    display: block;

    .list-item-owner,
    .list-item-time,
    .list-item-progress {
      width: 100%;
      margin: 8px 0 0 0;
      display: flex;
      align-items: center;
    }

    .list-item-title {
      margin: 0;
      width: 80px;
    }

    .list-item-tools {
      margin-top: 8px;
      justify-content: flex-end;
    }
  }
}
</style>
