<template>
  <div>
    <ele-card style="border-radius: 0">
      <ele-text type="heading" size="lg">Edit user</ele-text>
      <ele-text type="placeholder" style="margin-top: 6px">
        Modify user basic information after clicking the save button
      </ele-text>
    </ele-card>
    <ele-page>
      <ele-card>
        <ele-loading :loading="loading">
          <edit-form v-if="user" :data="user" />
        </ele-loading>
      </ele-card>
    </ele-page>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import { EleMessage } from 'ele-admin-plus/es';
import { usePageTab } from '@/utils/use-page-tab';
import EditForm from '../components/edit-form.vue';
import { getUser } from '@/api/system/user';

const { setPageTabTitle } = usePageTab();
const route = useRoute();
const userId = Number(route.params.id);

/** 查询状态 */
const loading = ref(true);

/** 用户信息 */
const user = ref();

/** 查询用户信息 */
const query = () => {
  if (!userId || isNaN(userId)) {
    return;
  }
  getUser(userId)
    .then((data) => {
      loading.value = false;
      user.value = data;
      // 修改页签标题
      setPageTabTitle(`Edit user[${user.value.nickname}]`);
    })
    .catch((e) => {
      EleMessage.error(e.message);
    });
};

// query();
</script>

<script>
export default {
  name: 'ListBasicEdit'
};
</script>
