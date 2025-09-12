<!-- 统计卡片 -->
<template>
  <el-row :gutter="16">
    <el-col :md="6" :sm="12" :xs="24">
      <ele-card class="statistics-card">
        <ele-text type="placeholder" class="statistics-header">
          <div class="statistics-header-text"></div>
          <ele-tooltip content="Indicator Description" placement="top" :offset="6">
            <el-icon class="statistics-header-tip">
              <QuestionCircleOutlined />
            </el-icon>
          </ele-tooltip>
        </ele-text>
        <ele-text size="sm" style="color:#6E7079" strong class="statistics-value">Register today 100</ele-text>
        <ele-text size="sm" style="color:#6E7079" strong class="statistics-value">First deposit number 50</ele-text>
        <ele-text size="sm" style="color:#6E7079" strong class="statistics-value">First deposit amount 300</ele-text>
        <ele-text size="sm" style="color:#6E7079" strong class="statistics-value">Total deposit amount 1000</ele-text>
        <ele-text size="sm" style="color:#6E7079" strong class="statistics-value"></ele-text>
        <el-divider />
      </ele-card>
    </el-col>
    <el-col :md="6" :sm="12" :xs="24">
      <ele-card class="statistics-card">
        <ele-text type="placeholder" class="statistics-header">
          <div class="statistics-header-text">Visits</div>
          <el-tag type="danger" size="small" :disable-transitions="true">
            Day
          </el-tag>
        </ele-text>
        <ele-text size="sm" style="color:#6E7079" strong class="statistics-value">
          <div class="statistics-header-text">Today's deposit</div>
          <span>100</span>
        </ele-text>
        <ele-text size="sm" style="color:#6E7079" strong class="statistics-value">
          <div class="statistics-header-text">Deposit number</div>
          <span>100</span>
        </ele-text>
        <ele-text size="sm" style="color:#6E7079" strong class="statistics-value">
          <div class="statistics-header-text">Deposit number</div>
          <span>100</span>
        </ele-text>
        <ele-text size="sm" style="color:#6E7079" strong class="statistics-value">
          <div class="statistics-header-text">Average deposit</div>
          <span>100</span>
        </ele-text>
        <!--        <v-chart-->
        <!--          ref="visitChartRef"-->
        <!--          :option="visitChartOption"-->
        <!--          style="height: 36px"-->
        <!--        />-->
        <el-divider />
      </ele-card>
    </el-col>
    <el-col :md="6" :sm="12" :xs="24">
      <ele-card class="statistics-card">
        <ele-text type="placeholder" class="statistics-header">
          <div class="statistics-header-text">Pay number</div>
          <el-tag size="small" :disable-transitions="true">Month</el-tag>
        </ele-text>
        <ele-text size="xxl" class="statistics-value">6,560</ele-text>
        <v-chart ref="payNumChartRef" :option="payNumChartOption" style="height: 36px" />
        <el-divider />
        <div>Conversion rate 60%</div>
      </ele-card>
    </el-col>
    <el-col :md="6" :sm="12" :xs="24">
      <ele-card class="statistics-card">
        <ele-text type="placeholder" class="statistics-header">
          <div class="statistics-header-text">Pay number</div>
          <el-tag type="success" size="small" :disable-transitions="true">
            Week
          </el-tag>
        </ele-text>
        <ele-text size="xxl" class="statistics-value">78%</ele-text>
        <div class="statistics-body">
          <el-progress color="#13c2c2" :percentage="78" :show-text="false" :stroke-width="8" style="width: 100%" />
        </div>
        <el-divider />
        <div class="statistics-footer">
          <div class="statistics-trend-text">
            <div>Weekly comparison 12% </div>
            <ele-text type="danger" :icon="CaretUpFilled" />
          </div>
          <div class="statistics-trend-text">
            <div>Daily comparison 11% </div>
            <ele-text type="success" :icon="CaretDownFilled" />
          </div>
        </div>
      </ele-card>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { use } from 'echarts/core';
import { CanvasRenderer } from 'echarts/renderers';
import { LineChart, BarChart } from 'echarts/charts';
import { GridComponent, TooltipComponent } from 'echarts/components';
import VChart from 'vue-echarts';
import {
  QuestionCircleOutlined,
  CaretUpFilled,
  CaretDownFilled
} from '@/components/icons';
import { getPayNumList } from '@/api/dashboard/analysis';
import { useEcharts } from '@/utils/use-echarts';

use([CanvasRenderer, LineChart, BarChart, GridComponent, TooltipComponent]);

/** Visits chart */
const visitChartRef = ref(null);

/** Pay number chart */
const payNumChartRef = ref(null);

useEcharts([visitChartRef, payNumChartRef]);

/** Visits chart option */
const visitChartOption = reactive({});

/** Pay number chart option */
const payNumChartOption = reactive({});

/** Get pay number data */
const getPayNumData = () => {
  getPayNumList()
    .then((data) => {
      Object.assign(visitChartOption, {
        color: '#975fe5',
        tooltip: {
          trigger: 'axis',
          formatter: `<i style="background: #975fe5;
                                  width: 10px;
                                  height: 10px;
                                  margin-right: 5px;
                                  border-radius: 50%;
                                  display: inline-block;"
                        ></i>{b0}: {c0}`
        },
        grid: {
          top: 0,
          bottom: 0,
          left: 0,
          right: 0
        },
        xAxis: [
          {
            show: false,
            type: 'category',
            boundaryGap: false,
            data: data.map((d) => d.date)
          }
        ],
        yAxis: [
          {
            show: false,
            type: 'value',
            splitLine: {
              show: false
            }
          }
        ],
        series: [
          {
            type: 'line',
            smooth: true,
            symbol: 'none',
            areaStyle: {
              opacity: 0.5
            },
            data: data.map((d) => d.value)
          }
        ]
      });

      Object.assign(payNumChartOption, {
        tooltip: {
          trigger: 'axis',
          formatter: `<i style="background: #5b8ff9;
                                  width: 10px;
                                  height: 10px;
                                  margin-right: 5px;
                                  border-radius: 50%;
                                  display: inline-block;"
                        ></i>{b0}: {c0}`
        },
        grid: {
          top: 0,
          bottom: 0,
          left: 0,
          right: 0
        },
        xAxis: [
          {
            show: false,
            type: 'category',
            data: data.map((d) => d.date)
          }
        ],
        yAxis: [
          {
            show: false,
            type: 'value',
            splitLine: {
              show: false
            }
          }
        ],
        series: [
          {
            type: 'bar',
            data: data.map((d) => d.value)
          }
        ]
      });
    })
    .catch((e) => {
      EleMessage.error(e.message);
    });
};

// getPayNumData();
</script>

<style lang="scss" scoped>
.statistics-card {
  position: relative;

  :deep(.ele-card-body) {
    padding: 16px 33px 12px 33px;
  }

  :deep(.el-divider) {
    margin: 12px 0;
    opacity: 0.6;
  }

  .statistics-header {
    display: flex;
    align-items: center;

    .statistics-header-text {
      flex: 1;
    }

    .statistics-header-tip {
      font-size: 15px;
      cursor: help;
    }
  }

  .statistics-value {
    display: flex;
    margin-top: 10px;
  }

  .statistics-body {
    height: 36px;
    display: flex;
    padding-top: 18px;
    box-sizing: border-box;
  }

  .statistics-trend-text {
    display: flex;
    align-items: center;
    white-space: nowrap;
    word-break: break-all;
    overflow: hidden;

    .el-icon {
      font-size: 16px;
      margin-left: 4px;
    }

    &+.statistics-trend-text {
      margin-left: 14px;
    }
  }

  .statistics-footer {
    display: flex;
    align-items: center;
  }
}
</style>
