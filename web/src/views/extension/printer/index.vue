<template>
  <ele-page multi-card>
    <ele-card header="Printer">
      <el-form label-width="80px" style="max-width: 320px" @submit.prevent="">
        <el-form-item label="Paper direction">
          <el-select clearable class="ele-fluid" placeholder="Not set" v-model="option.direction">
            <el-option value="landscape" label="Landscape" />
            <el-option value="portrait" label="Portrait" />
          </el-select>
        </el-form-item>
        <el-form-item label="Paper rotation">
          <el-select clearable class="ele-fluid" placeholder="Not rotate" v-model="option.orientation">
            <el-option value="rotate-left" label="Rotate left" />
            <el-option value="rotate-right" label="Rotate right" />
          </el-select>
        </el-form-item>
        <el-form-item label="Page margin">
          <el-input clearable v-model="option.margin" placeholder="Not set" />
        </el-form-item>
        <el-form-item label="Page title">
          <el-input clearable v-model="option.title" placeholder="Default" />
        </el-form-item>
        <el-form-item label="Print position">
          <el-select class="ele-fluid" v-model="option.target">
            <el-option value="_iframe" label="Sub window" />
            <el-option value="_self" label="Current window" />
          </el-select>
        </el-form-item>
        <el-form-item label="Always show">
          <el-radio-group v-model="option.static">
            <el-radio :value="false" label="No" />
            <el-radio :value="true" label="Yes" />
          </el-radio-group>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onPrint">Print</el-button>
        </el-form-item>
      </el-form>
      <ele-printer v-model="printing" :direction="option.direction" :orientation="option.orientation"
        :margin="option.margin" :title="option.title" :target="option.target" :static="option.static"
        :body-style="{ overflow: 'hidden' }" @done="onPrintDone">
        <div style="overflow: auto">
          <ele-table has-footer style="min-width: 600px" size="large">
            <colgroup>
              <col width="38px" />
              <col width="200px" />
              <col width="140px" />
              <col width="138px" />
              <col width="138px" />
              <col width="118px" />
              <col width="180px" />
            </colgroup>
            <thead>
              <tr>
                <th></th>
                <th>Project name</th>
                <th>Amount (10,000 yuan)</th>
                <th>Start time</th>
                <th>End time</th>
                <th>Status</th>
                <th>Progress</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(row, index) in projectList" :key="row.id">
                <td style="text-align: center">{{ index + 1 }}</td>
                <td>{{ row.projectName }}</td>
                <td>
                  <el-input-number :min="0" :max="999" v-model="row.money" placeholder="Input" controls-position="right"
                    class="ele-fluid" />
                </td>
                <td>{{ row.startDate }}</td>
                <td>{{ row.endDate }}</td>
                <td>
                  <ele-text v-if="row.status === 0" type="success">
                    In progress
                  </ele-text>
                  <ele-text v-else-if="row.status === 1" type="danger">
                    Delayed
                  </ele-text>
                  <ele-text v-else-if="row.status === 2" type="warning">
                    Not started
                  </ele-text>
                  <ele-text v-else-if="row.status === 3" type="info">
                    Ended
                  </ele-text>
                </td>
                <td>
                  <el-progress :percentage="row.progress" />
                </td>
              </tr>
            </tbody>
            <tfoot>
              <tr>
                <td></td>
                <td>Total</td>
                <td>{{ sumMoney }} 10,000 yuan</td>
                <td></td>
                <td></td>
                <td></td>
                <td></td>
              </tr>
            </tfoot>
          </ele-table>
        </div>
        <div style="margin: 12px 0">More form components demo：</div>
        <div>
          <el-radio-group v-model="status">
            <el-radio value="Unfinished" label="Unfinished" />
            <el-radio value="Completed" label="Completed" />
          </el-radio-group>
        </div>
        <div>
          <el-checkbox-group v-model="citys">
            <el-checkbox value="Wuhan" label="Wuhan" />
            <el-checkbox value="Beijing" label="Beijing" />
            <el-checkbox value="Shanghai" label="Shanghai" />
            <el-checkbox value="Guangzhou" label="Guangzhou" />
            <el-checkbox value="Shenzhen" label="Shenzhen" />
          </el-checkbox-group>
        </div>
        <div style="margin: 12px 0">Customize a checkbox component that is more suitable for printing style：</div>
        <div>
          <print-checkbox v-model="citys" :items="[
            { label: 'Wuhan', value: 'Wuhan' },
            { label: 'Beijing', value: 'Beijing' },
            { label: 'Shanghai', value: 'Shanghai' },
            { label: 'Guangzhou', value: 'Guangzhou' },
            { label: 'Shenzhen', value: 'Shenzhen' }
          ]" />
        </div>
      </ele-printer>
    </ele-card>
    <ele-card header="Use Cases">
      <el-form label-width="80px" style="max-width: 320px" @submit.prevent="">
        <el-form-item label="Party A">
          <el-input clearable v-model="form.partyA" />
        </el-form-item>
        <el-form-item label="Party B">
          <el-input clearable v-model="form.partyB" />
        </el-form-item>
        <el-form-item label="Address">
          <el-input clearable v-model="form.address" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onPrint2">Print</el-button>
        </el-form-item>
      </el-form>
      <print-contract ref="contractRef" :data="contractData" :is-static="option.static" :target="option.target" />
    </ele-card>
    <ele-card header="More Features">
      <el-button type="primary" @click="onPrintPdf">Print pdf</el-button>
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref, reactive, computed, nextTick } from 'vue';
import { printPdf } from 'ele-admin-plus/es/ele-printer/util';
import { EleMessage } from 'ele-admin-plus/es';
import { getProjectList } from '@/api/example';
import PrintCheckbox from './components/print-checkbox.vue';
import PrintContract from './components/print-contract.vue';
import dayjs from 'dayjs';

/** Print */
const printing = ref(false);

/** Printing parameters */
const option = reactive({
  direction: void 0,
  orientation: void 0,
  margin: '',
  title: '',
  target: '_iframe',
  static: false
});

/** Print */
const onPrint = () => {
  console.log('开始打印...');
  printing.value = true;
};

/** Print done */
const onPrintDone = () => {
  console.log('Print done.');
};

/** Status */
const status = ref('Unfinished');

/** City */
const citys = ref(['Wuhan', 'Beijing']);

/** Project progress data */
const projectList = ref([]);

/** Amount total */
const sumMoney = computed(() => {
  return projectList.value.reduce((sum, d) => {
    const money = d.money ? Number(d.money) : 0;
    return sum + (isNaN(money) ? 0 : money);
  }, 0);
});

/** Query project progress */
const queryProjectList = () => {
  getProjectList()
    .then((data) => {
      projectList.value = data;
    })
    .catch((e) => {
      console.error(e);
    });
};

queryProjectList();

/** Form data */
const form = reactive({
  partyA: 'XX House Rental Company',
  partyB: 'Zhang San',
  address: 'Room 1104, Unit 1, Building 18, XX Community, XX Street, XX District, XX City, XX Province'
});

/** Contract component */
const contractRef = ref(null);

/** Contract data */
const contractData = ref(null);

/** Print contract */
const onPrint2 = () => {
  contractData.value = {
    ...form,
    date: dayjs().format('YYYY-MM-DD'),
    projects: [
      ...projectList.value,
      ...projectList.value,
      ...projectList.value,
      ...projectList.value
    ].map((d, i) => {
      return { ...d, id: d.id + '_' + i, projectName: d.projectName + i };
    })
  };
  nextTick(() => {
    contractRef.value && contractRef.value.print();
  });
};

/** 打印pdf */
const onPrintPdf = () => {
  const loading = EleMessage.loading({
    message: 'Loading ...',
    plain: true,
    centered: true,
    mask: true
  });
  printPdf({
    url: 'https://cdn.eleadmin.com/20200610/20200708224450.pdf',
    done: () => {
      loading.close();
    }
  });
};
</script>

<script>
export default {
  name: 'ExtensionPrinter'
};
</script>
