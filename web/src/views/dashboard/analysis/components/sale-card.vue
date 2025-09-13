<template>
  <ele-card header="Online Users" :header-style="{ paddingTop: 5, paddingBottom: 5 }" :body-style="{ padding: 0 }">
    <template #extra>
      <div class="hidden-xs-only" style="display: flex; align-items: center">
        <el-radio-group>
          <el-radio value="small" label="Today" />
          <el-radio value="default" label="Yesterday" />
        </el-radio-group>
        <el-radio-group style="margin-left: 12px" v-model="saleSearch.dateType" @change="onDateTypeChange">
          <el-radio-button value="1" label="Online Users" />
          <el-radio-button value="2" label="Yesterday" />
          <el-radio-button value="3" label="Today" />
          <el-radio-button value="4" label="Total" />
          <el-radio-button value="5" label="Total" />
          <el-radio-button value="6" label="Total" />
        </el-radio-group>

        <div class="hidden-md-and-down" style="width: 320px; margin-left: 12px">
          <el-date-picker unlink-panels type="datetimerange" v-model="saleSearch.datetime" range-separator="-"
            format="YYYY-MM-DD HH:mm" value-format="YYYY-MM-DD HH:mm" start-placeholder="Start Time"
            end-placeholder="End Time" class="ele-fluid" />
        </div>
      </div>
    </template>
    <div class="sale-body">
      <el-row :gutter="16">
        <el-col :md="7" :sm="9" :xs="24">
          <div v-if="saleSearch.type === 'saleroom'" style="padding: 0 0 8px 0">
            <div class="sale-body-title">Online Users</div>
            <div v-for="(item, index) in saleroomRankData" :key="index" class="sale-rank-item">
              <!--              <el-tag-->
              <!--                size="small"-->
              <!--                :disable-transitions="true"-->
              <!--                :type="index < 3 ? void 0 : 'info'"-->
              <!--                :effect="index < 3 ? 'dark' : 'light'"-->
              <!--                :color="index < 3 ? '#314659' : void 0"-->
              <!--                style="border: none; border-radius: 50%; width: 20px"-->
              <!--              >-->
              <!--                {{ index + 1 }}-->
              <!--              </el-tag>-->
              <ele-ellipsis class="sale-rank-item-text">
                {{ item.name }}
              </ele-ellipsis>
              <ele-text type="placeholder">{{ item.value }}</ele-text>
            </div>
          </div>
          <div v-else style="padding: 0 0 8px 0">
            <div class="sale-body-title">Online Users</div>
            <div v-for="(item, index) in visitsRankData" :key="index" class="sale-rank-item">
              <el-tag size="small" :disable-transitions="true" :type="index < 3 ? void 0 : 'info'"
                :effect="index < 3 ? 'dark' : 'light'" :color="index < 3 ? '#314659' : void 0"
                style="border: none; border-radius: 50%; width: 20px">
                {{ index + 1 }}
              </el-tag>
              <ele-ellipsis class="sale-rank-item-text">
                {{ item.name }}
              </ele-ellipsis>
              <ele-text type="placeholder">{{ item.value }}</ele-text>
            </div>
          </div>
        </el-col>
        <el-col :md="17" :sm="15" :xs="24">
          <div class="sale-body-title">
            {{ saleSearch.type === 'saleroom' ? 'Online Users Trend' : 'Visits Trend' }}
          </div>
          <v-chart ref="saleChartRef" :option="saleChartOption" style="height: 320px" @click="onSaleChartClick" />
        </el-col>
      </el-row>
    </div>
  </ele-card>
</template>

<script setup>
import { ref, reactive } from 'vue';
import dayjs from 'dayjs';
import { EleMessage } from 'ele-admin-plus/es';
import { use } from 'echarts/core';
import { CanvasRenderer } from 'echarts/renderers';
import { BarChart } from 'echarts/charts';
import { GridComponent, TooltipComponent } from 'echarts/components';
import VChart from 'vue-echarts';
import { getSaleroomList } from '@/api/dashboard/analysis';
import { useEcharts } from '@/utils/use-echarts';

use([CanvasRenderer, BarChart, GridComponent, TooltipComponent]);

const saleChartRef = ref(null);

useEcharts([saleChartRef]);

/** Sales bar chart configuration */
const saleChartOption = reactive({});

/** Online users rank data */
const saleroomRankData = ref([
  { name: 'Current number of online users', value: '333' },
  { name: 'top up', value: '333' },
  { name: 'Withdrawal', value: '333' },
  { name: 'register', value: '333' },
  { name: 'Platform profit and loss', value: '333' },
  { name: 'Coding on the day', value: '333' }
]);

/** Visits rank data */
const visitsRankData = ref([
  { name: 'Gongzhuan Road Store No. 1', value: '222,001' },
  { name: 'Gongzhuan Road Store No. 2', value: '222,002' },
  { name: 'Gongzhuan Road Store No. 3', value: '222,002' },
  { name: 'Gongzhuan Road Store No. 4', value: '222,004' },
  { name: 'Gongzhuan Road Store No. 5', value: '222,005' },
  { name: 'Gongzhuan Road Store No. 6', value: '222,006' },
  { name: 'Gongzhuan Road Store No. 7', value: '222,007' }
]);

/** Sales volume trend data */
const saleroomData1 = ref([]);

/** Visits trend data */
const saleroomData2 = ref([]);

/** Sales volume search parameters */
const saleSearch = reactive({
  type: 'saleroom',
  dateType: '2',
  datetime: void 0
});

/** Get sales volume data */
const getSaleroomData = () => {
  getSaleroomList()
    .then((data) => {
      saleroomData1.value = data.list1;
      saleroomData2.value = data.list2;
      onSaleTypeChange();
    })
    .catch((e) => {
      EleMessage.error(e.message);
    });
};

/** Sales volume tab selection change event */
const onSaleTypeChange = () => {
  if (saleSearch.type === 'saleroom') {
    Object.assign(saleChartOption, {
      tooltip: {
        trigger: 'axis'
      },
      xAxis: [
        {
          type: 'category',
          data: saleroomData1.value.map((d) => d.month)
        }
      ],
      yAxis: [
        {
          type: 'value'
        }
      ],
      series: [
        {
          type: 'bar',
          data: saleroomData1.value.map((d) => d.value)
        }
      ]
    });
  } else {
    Object.assign(saleChartOption, {
      tooltip: {
        trigger: 'axis'
      },
      xAxis: [
        {
          type: 'category',
          data: saleroomData2.value.map((d) => d.month)
        }
      ],
      yAxis: [
        {
          type: 'value'
        }
      ],
      series: [
        {
          type: 'bar',
          data: saleroomData2.value.map((d) => d.value)
        }
      ]
    });
  }
};

/** Sales volume chart click event */
const onSaleChartClick = (params) => {
  const type = saleSearch.type === 'saleroom' ? 'Sales volume' : 'Visits';
  EleMessage.info({
    message: `Clicked: ${params.name} , ${type}: ${params.value}`,
    plain: true
  });
};

/** Date radio switch event */
const onDateTypeChange = (dateType) => {
  const type = ['day', 'week', 'month', 'year'][Number(dateType) - 1];
  const now = dayjs();
  const formatter = 'YYYY-MM-DD HH:mm';
  const start = now.startOf(type).format(formatter);
  const end = now.endOf(type).format(formatter);
  saleSearch.datetime = [start, end];
};

onDateTypeChange('2');

getSaleroomData();
</script>

<style lang="scss" scoped>
.sale-body {
  padding: 16px 0 10px 0;
}

.sale-body-title {
  padding: 6px 20px;
}

.sale-rank-item {
  display: flex;
  align-items: center;
  padding: 0 20px;
  margin-top: 18px;
  box-sizing: border-box;

  .sale-rank-item-text {
    flex: 1;
    padding-left: 12px;
  }
}
</style>
