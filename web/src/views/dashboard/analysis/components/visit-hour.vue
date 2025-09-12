<template>
  <ele-card header="Visits in the last hour" :body-style="{ padding: '16px 6px 0 0' }">
    <v-chart ref="visitHourChartRef" :option="visitHourChartOption" style="height: 362px" />
  </ele-card>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { use } from 'echarts/core';
import { CanvasRenderer } from 'echarts/renderers';
import { LineChart } from 'echarts/charts';
import {
  GridComponent,
  TooltipComponent,
  LegendComponent
} from 'echarts/components';
import VChart from 'vue-echarts';
import { getVisitHourList } from '@/api/dashboard/analysis';
import { useEcharts } from '@/utils/use-echarts';

use([
  CanvasRenderer,
  LineChart,
  GridComponent,
  TooltipComponent,
  LegendComponent
]);

const visitHourChartRef = ref(null);

useEcharts([visitHourChartRef]);

/** Configuration of the line graph of visits in the last hour */
const visitHourChartOption = reactive({});

/** Get access data for the last hour */
const getVisitHourData = () => {
  getVisitHourList()
    .then((data) => {
      Object.assign(visitHourChartOption, {
        tooltip: {
          trigger: 'axis'
        },
        legend: {
          data: ['Visits', 'Accesses'],
          right: 20
        },
        xAxis: [
          {
            type: 'category',
            boundaryGap: false,
            data: data.map((d) => d.time)
          }
        ],
        yAxis: [
          {
            type: 'value'
          }
        ],
        series: [
          {
            name: 'Visits',
            type: 'line',
            smooth: true,
            symbol: 'none',
            areaStyle: {
              opacity: 0.5
            },
            data: data.map((d) => d.views)
          }
        ]
      });
    })
    .catch((e) => {
      EleMessage.error(e.message);
    });
};

getVisitHourData();
</script>
