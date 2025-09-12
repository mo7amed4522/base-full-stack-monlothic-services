<template>
  <ele-card header="Online number">
    <div style="text-align: center">
      <ele-text type="placeholder">{{ currentTime }}</ele-text>
      <ele-text size="xxxl">
        <ele-count-up :end-val="onlineNum" />
      </ele-text>
      <ele-text type="placeholder" style="margin-bottom: 11px">
        Total online number
      </ele-text>
      <ele-dot :text="updateTimeText" size="8px" />
    </div>
  </ele-card>
</template>

<script setup>
import { ref, computed, onBeforeUnmount } from 'vue';
import dayjs from 'dayjs';

/** Online number update timer */
let onlineNumTimer = null;

/** Total online number countdown */
const updateTime = ref(10);

/** Current time */
const currentTime = ref(dayjs().format('HH:mm:ss'));

/** Online number */
const onlineNum = ref(228);

/** Online number countdown text */
const updateTimeText = computed(() => updateTime.value + ' seconds later');

/** Online number update countdown */
const startUpdateOnlineNum = () => {
  onlineNumTimer = window.setInterval(() => {
    currentTime.value = dayjs().format('HH:mm:ss');
    if (updateTime.value === 1) {
      updateTime.value = 10;
      onlineNum.value = 100 + Math.round(Math.random() * 900);
    } else {
      updateTime.value--;
    }
  }, 1000);
};

onBeforeUnmount(() => {
  // Destroy timer
  if (onlineNumTimer) {
    clearInterval(onlineNumTimer);
    onlineNumTimer = null;
  }
});

startUpdateOnlineNum();
</script>
