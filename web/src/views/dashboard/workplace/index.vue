<template>
  <ele-page class="workplace-page">
    <profile-card />
    <link-card ref="linkCardRef" />
    <el-row :gutter="16" ref="wrapRef">
      <el-col v-for="(item, index) in data" :key="item.name" :md="item.md" :sm="item.sm" :xs="item.xs">
        <component :is="item.name" :title="item.title" @command="(command) => onCommand(command, index)" />
      </el-col>
    </el-row>
    <ele-card :body-style="{ padding: 0 }">
      <div class="workplace-bottom">
        <el-link type="primary" :icon="PlusCircleOutlined" :underline="false" class="workplace-button" @click="add">
          Add card
        </el-link>
        <el-divider direction="vertical" style="margin: 0" />
        <el-link type="primary" :icon="UndoOutlined" :underline="false" class="workplace-button" @click="reset">
          Reset layout
        </el-link>
      </div>
    </ele-card>
    <!-- Add a card pop-up -->
    <ele-modal :width="680" v-model="visible" title="Add a card">
      <el-row :gutter="16" style="margin-top: -16px">
        <el-col v-for="item in notAddedData" :key="item.name" :md="8" :sm="12" :xs="24">
          <ele-card bordered :header="item.title" :header-style="{ padding: '8px 14px', fontSize: '14px' }"
            :body-style="{ padding: '32px 0', textAlign: 'center' }" style="margin-top: 16px">
            <el-button plain round size="small" type="primary" @click="addView(item)">
              Add
            </el-button>
          </ele-card>
        </el-col>
      </el-row>
      <el-empty v-if="!notAddedData.length" description="All views have been added" :image-size="80" />
    </ele-modal>
  </ele-page>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue';
import SortableJs from 'sortablejs';
import { EleMessage } from 'ele-admin-plus/es';
import { PlusCircleOutlined, UndoOutlined } from '@/components/icons';
import ProfileCard from './components/profile-card.vue';
import LinkCard from './components/link-card.vue';
const CACHE_KEY = 'workplace-layout';

/** Default layout */
const DEFAULT = [
  {
    name: 'activities-card',
    title: 'Latest activities',
    md: 8,
    sm: 24,
    xs: 24
  },
  {
    name: 'task-card',
    title: 'My tasks',
    md: 8,
    sm: 24,
    xs: 24
  },
  {
    name: 'goal-card',
    title: 'Monthly goals',
    md: 8,
    sm: 24,
    xs: 24
  },
  {
    name: 'project-card',
    title: 'Project progress',
    md: 16,
    sm: 24,
    xs: 24
  },
  {
    name: 'user-list',
    title: 'Group members',
    md: 8,
    sm: 24,
    xs: 24
  }
];

/** Get cached data */
const getCacheData = () => {
  const str = localStorage.getItem(CACHE_KEY);
  try {
    return str ? JSON.parse(str) : null;
  } catch (e) {
    return null;
  }
};

/** Card data */
const data = ref([...(getCacheData() ?? DEFAULT)]);

/** Modal is open */
const visible = ref(false);

/** Shortcut card */
const linkCardRef = ref(null);

/** Container */
const wrapRef = ref(null);

/** Drag sort instance */
let sortableIns = null;

/** Unadded data */
const notAddedData = computed(() => {
  return DEFAULT.filter((d) => !data.value.some((t) => t.name === d.name));
});

/** Add */
const add = () => {
  visible.value = true;
};

/** Reset layout */
const reset = () => {
  data.value = [...DEFAULT];
  cacheData();
  linkCardRef.value?.reset?.();
  EleMessage.success('Layout reset');
};

/** Cache layout */
const cacheData = () => {
  localStorage.setItem(CACHE_KEY, JSON.stringify(data.value));
};

/** Edit card */
const onCommand = (command, index) => {
  switch (command) {
    case 'edit': // Edit
      EleMessage.info({ message: 'Click edit', plain: true });
      break;
    case 'remove': // Remove
      data.value = data.value.filter((_d, i) => i !== index);
      cacheData();
      break;
  }
};

/** Add view */
const addView = (item) => {
  data.value.push(item);
  cacheData();
  EleMessage.success('Add view');
};

onMounted(() => {
  // Card support drag sort
  if ('ontouchstart' in document.documentElement) {
    return;
  }
  sortableIns = new SortableJs(wrapRef.value?.$el, {
    handle: '.ele-card-header',
    animation: 300,
    onUpdate: ({ oldIndex, newIndex }) => {
      if (typeof oldIndex === 'number' && typeof newIndex === 'number') {
        const temp = [...data.value];
        temp.splice(newIndex, 0, temp.splice(oldIndex, 1)[0]);
        data.value = temp;
        cacheData();
      }
    },
    setData: () => { },
    forceFallback: true
  });
});

onBeforeUnmount(() => {
  sortableIns && sortableIns.destroy();
  sortableIns = null;
});
</script>

<script>
import ActivitiesCard from './components/activities-card.vue';
import TaskCard from './components/task-card.vue';
import GoalCard from './components/goal-card.vue';
import ProjectCard from './components/project-card.vue';
import UserList from './components/user-list.vue';

export default {
  name: 'DashboardWorkplace',
  components: {
    ActivitiesCard,
    TaskCard,
    GoalCard,
    ProjectCard,
    UserList
  }
};
</script>

<style lang="scss" scoped>
.workplace-page {
  :deep(.ele-card-header) {
    cursor: move;
  }

  :deep(.el-col) {
    &.sortable-chosen>.ele-card {
      box-shadow: 0 2px 16px 0 rgba(0, 0, 0, 0.2);
    }

    &.sortable-ghost {
      opacity: 0;
    }

    &.sortable-fallback {
      opacity: 1 !important;
    }
  }
}

/* Bottom buttons */
.workplace-bottom {
  display: flex;
  align-items: center;

  .workplace-button {
    flex: 1;
    padding: 10px 0;
    transition: background-color 0.2s;

    &:hover {
      background: hsla(0, 0%, 60%, 0.05);
    }

    :deep(.el-icon) {
      font-size: 15px;
      margin: -1px 6px 0 0;
    }
  }
}
</style>
