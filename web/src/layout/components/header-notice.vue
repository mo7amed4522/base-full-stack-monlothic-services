<!-- Header notice -->
<template>
  <ele-popover :width="336" trigger="click" transition="el-zoom-in-top" :content-style="{ padding: 0 }"
    :body-style="{ overflow: 'hidden' }" :popper-style="{ maxWidth: 'calc(100% - 32px)' }" :popper-options="{
      strategy: 'fixed',
      modifiers: [{ name: 'offset', options: { offset: [0, 5] } }]
    }">
    <template #reference>
      <div style="display: flex; align-items: center; height: 100%">
        <el-badge :value="unreadNum" :hidden="!unreadNum" style="line-height: 1; padding: 4px 0">
          <el-icon style="transform: scale(1.17) translateY(1px)">
            <BellOutlined />
          </el-icon>
        </el-badge>
      </div>
    </template>
    <ele-tabs ref="tabRef" :center="true" v-model="active" :items="[
      { name: 'notice', label: 'Notice' },
      { name: 'letter', label: 'Letter' },
      { name: 'todo', label: 'Todo' }
    ]" class="notice-tabs">
      <template #label="{ item, label }">
        <span>{{ label }}</span>
        <span v-if="item.name === 'notice' && notice.length">({{ notice.length }})</span>
        <span v-if="item.name === 'letter' && letter.length">({{ letter.length }})</span>
        <span v-if="item.name === 'todo' && todo.length">({{ todo.length }})</span>
      </template>
      <template #notice>
        <div class="list-wrapper">
          <div v-for="item in notice" :key="item.id" class="list-item">
            <div class="list-item-icon" :style="{ background: item.color }">
              <el-icon>
                <Comment v-if="item.icon === 'report'" style="transform: translateY(1px)" />
                <Stamp v-else-if="item.icon === 'leave'" />
                <Promotion v-else-if="item.icon === 'meeting'" />
                <Checked v-else-if="item.icon === 'interview'" />
                <BellFilled v-else-if="item.icon === 'salary'" />
              </el-icon>
            </div>
            <div class="list-item-body">
              <ele-ellipsis>{{ item.title }}</ele-ellipsis>
              <ele-ellipsis type="placeholder" class="list-item-text">
                {{ item.time }}
              </ele-ellipsis>
            </div>
          </div>
        </div>
        <div v-if="notice.length" class="bottom-tools">
          <div class="bottom-tool" @click="clearNotice">Clear notice</div>
          <el-divider direction="vertical" style="margin: 0; width: 0" />
          <router-link to="/user/message?type=notice" class="bottom-tool">View more</router-link>
        </div>
        <el-empty v-else description="View all notice" :image-size="68" />
      </template>
      <template #letter>
        <div class="list-wrapper">
          <div v-for="item in letter" :key="item.id" class="list-item">
            <el-avatar :size="32" :src="item.avatar" />
            <div class="list-item-body">
              <ele-ellipsis>{{ item.title }}</ele-ellipsis>
              <ele-ellipsis type="placeholder" class="list-item-text">
                {{ item.content }}
              </ele-ellipsis>
              <ele-ellipsis type="placeholder" class="list-item-text">
                {{ item.time }}
              </ele-ellipsis>
            </div>
          </div>
        </div>
        <div v-if="letter.length" class="bottom-tools">
          <div class="bottom-tool" @click="clearLetter">Clear letter</div>
          <el-divider direction="vertical" style="margin: 0; width: 0" />
          <router-link to="/user/message?type=letter" class="bottom-tool">View more</router-link>
        </div>
        <el-empty v-else description="View all letter" :image-size="68" />
      </template>
      <template #todo>
        <div class="list-wrapper">
          <div v-for="item in todo" :key="item.id" class="list-item">
            <div class="list-item-body">
              <ele-ellipsis>{{ item.title }}</ele-ellipsis>
              <ele-ellipsis type="placeholder" class="list-item-text">
                {{ item.description }}
              </ele-ellipsis>
            </div>
            <el-tag v-if="item.status === 0" type="info" size="small" :disable-transitions="true">Not started</el-tag>
            <el-tag v-else-if="item.status === 1" type="danger" size="small" :disable-transitions="true">Expire
              soon</el-tag>
            <el-tag v-else-if="item.status === 2" size="small" :disable-transitions="true">In progress</el-tag>
          </div>
        </div>
        <div v-if="todo.length" class="bottom-tools">
          <div class="bottom-tool" @click="clearTodo">Clear To-Do</div>
          <el-divider direction="vertical" style="margin: 0; width: 0" />
          <router-link to="/user/message?type=todo" class="bottom-tool">See more</router-link>
        </div>
        <el-empty v-else description="View all todo" :image-size="68" />
      </template>
    </ele-tabs>
  </ele-popover>
</template>

<script setup>
import { ref, computed, nextTick } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { Comment, Stamp, Promotion, Checked, BellFilled } from '@element-plus/icons-vue';
import { BellOutlined } from '@/components/icons';
import { getUnreadNotice } from '@/api/example';

/** Tabs */
const tabRef = ref(null);

/** Tabs active */
const active = ref('notice');

/** Notice data */
const notice = ref([]);

/** Letter data */
const letter = ref([]);

/** Todo data */
const todo = ref([]);

/** Unread number */
const unreadNum = computed(() => {
  return notice.value.length + letter.value.length + todo.value.length;
});

/** Query data */
const query = () => {
  getUnreadNotice()
    .then((result) => {
      // notice.value = result.notice;
      // letter.value = result.letter;
      // todo.value = result.todo;
    })
    .catch((e) => {
      EleMessage.error(e.message);
    });
};

/** Clear notice */
const clearNotice = () => {
  notice.value = [];
  updateActiveBar();
};

/** Clear letter */
const clearLetter = () => {
  letter.value = [];
  updateActiveBar();
};

/** Clear todo */
const clearTodo = () => {
  todo.value = [];
  updateActiveBar();
};

/** Update active bar */
const updateActiveBar = () => {
  nextTick(() => {
    tabRef.value?.updateActiveBar?.();
  });
};

// query();
</script>

<style lang="scss" scoped>
.list-wrapper {
  padding-top: 8px;
  box-sizing: border-box;
  max-height: calc(100vh - 180px);
  min-height: 60px;
  overflow: auto;
}

.list-item {
  display: flex;
  padding: 14px 24px;
  box-sizing: border-box;
  transition: background-color 0.2s;
  cursor: pointer;

  .list-item-body {
    flex: 1;
    overflow: hidden;

    .list-item-text {
      margin-top: 6px;
    }
  }

  .list-item-icon {
    width: 32px;
    height: 32px;
    color: #fff;
    font-size: 16px;
    border-radius: 50%;
    text-align: center;
    background-color: #60b2fc;
    flex-shrink: 0;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  :deep(.el-avatar) {
    flex-shrink: 0;
  }

  .list-item-icon+.list-item-body,
  :deep(.el-avatar + .list-item-body) {
    padding-left: 12px;
  }

  &:hover {
    background-color: hsla(0, 0%, 60%, 0.08);
  }

  &+.list-item {
    border-top: 1px solid hsla(0, 0%, 60%, 0.2);
  }
}

/* 操作按钮 */
.bottom-tools {
  display: flex;
  align-items: center;
  border-top: 1px solid hsla(0, 0%, 60%, 0.2);

  .bottom-tool {
    flex: 1;
    line-height: 46px;
    text-align: center;
    text-decoration: none;
    transition: background-color 0.3s;
    cursor: pointer;
    color: inherit;

    &:hover {
      background: hsla(0, 0%, 60%, 0.08);
    }
  }
}

/* Modify tab style */
.notice-tabs :deep(.el-tabs__header) {
  --ele-tab-height: 44px;
  --ele-tab-padding: 22px;
}
</style>
