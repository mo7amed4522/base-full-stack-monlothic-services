<template>
  <ele-page>
    <ele-card>
      <el-form ref="settingFormRef" label-width="100px" @submit.prevent="">
        <el-row :gutter="8">
          <el-form-item label="Merchant Name：">
            <div class="title">万博</div>
          </el-form-item>
        </el-row>
        <el-row :gutter="8">
          <el-col :lg="24" :md="12" :sm="12" :xs="24">
            <el-form-item label="Gift Way：">
              <el-input clearable :disabled="isWbDisabled" v-model="wanBoForm.immediatePercentage"
                placeholder="Please input" style="width: 180px; margin-right: 20px">
                <!-- 前面追加 -->
                <template #prepend>immediate</template>
                <template #append>%</template>
              </el-input>
              <el-input clearable :disabled="isWbDisabled" v-model="wanBoForm.secondDayPercentage"
                placeholder="Please input" style="width: 180px">
                <template #prepend>second day</template>
                <template #append>%</template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="8">
          <el-col :lg="24" :md="12" :sm="12" :xs="24">
            <el-form-item>
              <el-button :type="isWbDisabled ? 'primary' : 'default'" @click="isWbDisabled = !isWbDisabled">{{
                isWbDisabled ? 'Edit' : 'Cancel' }}</el-button>
              <el-button :type="isWbDisabled ? 'default' : 'primary'" :disabled="isWbDisabled"
                @click="updateConfigUser('paymer01')">Confirm</el-button>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="8">
          <el-form-item label="Merchant name: ">
            <div class="title">Asia Investment</div>
          </el-form-item>
        </el-row>
        <el-row :gutter="8">
          <el-col :lg="24" :md="12" :sm="12" :xs="24">
            <el-form-item label="Gift Way：">
              <el-input clearable :disabled="isYtDisabled" v-model="yaTouForm.immediatePercentage"
                placeholder="Please input" style="width: 180px; margin-right: 20px">
                <!-- 前面追加 -->
                <template #prepend>immediate</template>
                <template #append>%</template>
              </el-input>
              <el-input clearable :disabled="isYtDisabled" v-model="yaTouForm.secondDayPercentage"
                placeholder="Please input" style="width: 180px">
                <template #prepend>second day</template>
                <template #append>%</template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="8">
          <el-col :lg="24" :md="12" :sm="12" :xs="24">
            <el-form-item>
              <el-button :type="isYtDisabled ? 'primary' : 'default'" @click="isYtDisabled = !isYtDisabled">{{
                isYtDisabled ? 'Edit' : 'Cancel' }}</el-button>
              <el-button :type="isYtDisabled ? 'default' : 'primary'" :disabled="isYtDisabled"
                @click="updateConfigUser('paymer02')">Confirm</el-button>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="8">
          <el-form-item label="Merchant name: ">
            <div class="title">Asia Investment</div>
          </el-form-item>
        </el-row>
        <el-row :gutter="8">
          <el-col :lg="24" :md="12" :sm="12" :xs="24">
            <el-form-item label="Gift Way：">
              <el-input clearable :disabled="isYgDisabled" v-model="yaGuanForm.immediatePercentage"
                placeholder="Please input" style="width: 180px; margin-right: 20px">
                <!-- 前面追加 -->
                <template #prepend>immediate</template>
                <template #append>%</template>
              </el-input>
              <el-input clearable :disabled="isYgDisabled" v-model="yaGuanForm.secondDayPercentage"
                placeholder="Please input" style="width: 180px">
                <template #prepend>second day</template>
                <template #append>%</template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="8">
          <el-col :lg="24" :md="12" :sm="12" :xs="24">
            <el-form-item>
              <el-button :type="isYgDisabled ? 'primary' : 'default'" @click="isYgDisabled = !isYgDisabled">{{
                isYgDisabled ? 'Edit' : 'Cancel' }}</el-button>
              <el-button :type="isYgDisabled ? 'default' : 'primary'" :disabled="isYgDisabled"
                @click="updateConfigUser('paymer03')">Confirm</el-button>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="8">
          <el-col :lg="24" :md="12" :sm="12" :xs="24">
            <div class="tips">
              <div class="tips-title">说明</div>
              <div class="tips-content">
                <div>1. If both options have values, then a portion of the immediate grant is sent immediately, and the
                  remaining proportion is sent on the specified time of the second day.</div>
                <div>2. If the grant ratio is 0, then the current option is not applied.</div>
              </div>
            </div>
          </el-col>
        </el-row>
      </el-form>
    </ele-card>
  </ele-page>
</template>
<script setup>
import { ref, onMounted } from 'vue';
import { EleMessage } from 'ele-admin-plus';
import { configMUserApi, getConfigMUserApi } from '@/api/merchant';
import { useMessageBox } from 'ele-admin-plus/es';

const messageBox = useMessageBox();

const isWbDisabled = ref(true);

const isYtDisabled = ref(true);

const isYgDisabled = ref(true);

const wanBoForm = ref({
  immediatePercentage: 0,
  secondDayPercentage: 0,
  merName: ''
});

const yaTouForm = ref({
  immediatePercentage: 0,
  secondDayPercentage: 0
});

const yaGuanForm = ref({
  immediatePercentage: 0,
  secondDayPercentage: 0,
  merName: ''
});

// 获取商户配置
const getConfigMUser = async () => {
  try {
    const res = await getConfigMUserApi();
    wanBoForm.value = {
      ...res.paymer01,
      immediatePercentage: res.paymer01.immediatePercentage * 100,
      secondDayPercentage: res.paymer01.secondDayPercentage * 100
    }
    yaTouForm.value = {
      ...res.paymer02,
      immediatePercentage: res.paymer02.immediatePercentage * 100,
      secondDayPercentage: res.paymer02.secondDayPercentage * 100
    }
    yaGuanForm.value = {
      ...res.paymer03,
      immediatePercentage: res.paymer03.immediatePercentage * 100,
      secondDayPercentage: res.paymer03.secondDayPercentage * 100
    };
  } catch (e) {
    EleMessage.error(e.message);
  }
};

const getFormData = (merName) => {
  if (merName === 'paymer01') {
    return wanBoForm.value;
  } else if (merName === 'paymer02') {
    return yaTouForm.value;
  } else if (merName === 'paymer03') {
    return yaGuanForm.value;
  }
  return null;
};

const setDisabledState = (merName) => {
  if (merName === 'paymer01') {
    isWbDisabled.value = true;
  } else if (merName === 'paymer02') {
    isYtDisabled.value = true;
  } else if (merName === 'paymer03') {
    isYgDisabled.value = true;
  }
}

const getMerName = (merName) => {
  if (merName === 'paymer01') {
    return 'Asia Investment';
  } else if (merName === 'paymer02') {
    return 'Asia Investment';
  } else if (merName === 'paymer03') {
    return 'Asia Investment';
  }
}

// 赠送方式配置
const updateConfigUser = async (merName) => {
  messageBox
    .confirm(`Confirm the modification ${getMerName(merName)} Is the merchant’s gift configuration?? `, 'Operation confirmation', {
      type: 'warning',
      draggable: true
    })
    .then(async () => {
      try {
        const formData = getFormData(merName);
        const res = await configMUserApi({
          immediatePercentage: Math.max(Number(formData.immediatePercentage) / 100, 0),
          secondDayPercentage: Math.max(Number(formData.secondDayPercentage) / 100, 0),
          merName: merName
        });

        // 处理 disabled 状态
        setDisabledState(merName);

        EleMessage.success(res);
        getConfigMUser();
      } catch (error) {
        EleMessage.error(error.message);
      }
    })
    .catch(() => { });
};

onMounted(() => {
  getConfigMUser();
});
</script>
<style lang="scss" scoped>
.title {
  font-size: 14px;
  font-weight: bold;
}

.tips {
  display: flex;
  font-size: 12px;
  color: #999;
  margin-bottom: 10px;

  &-title {
    margin-right: 5px;
  }
}
</style>
