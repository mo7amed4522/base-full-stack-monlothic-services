<!-- Contract -->
<template>
  <ele-printer ref="printerRef" margin="0mm 12mm 10mm 12mm" :header-style="{
    padding: '26px 0 2px 0',
    fontSize: '13px',
    borderBottom: '1px solid #666',
    marginBottom: '26px'
  }" :body-style="{ padding: '0 14px', fontSize: '17px', lineHeight: 2.5 }" :target="target" :static="isStatic">
    <template #header>
      <div style="color: #444">
        <img src="/favicon.ico" style="height: 15px; width: 15px; vertical-align: -2px" />
        <span> XXX Apartment Lease Contract</span>
      </div>
      <div style="color: #888">Phone：027-987654321</div>
    </template>
    <div v-if="!!data" style="page-break-after: always">
      <div style="font-size: 35px; text-align: center">Contract</div>
      <div style="margin-top: 12px">
        <span>Party A：</span>
        <span class="demo-input" style="min-width: 280px">
          {{ data.partyA }}
        </span>
      </div>
      <div>
        <span>Party B：</span>
        <span class="demo-input" style="min-width: 280px">
          {{ data.partyB }}
        </span>
      </div>
      <div>
        <span>Address：</span>
        <span class="demo-input" style="min-width: 280px">
          {{ data.address }}
        </span>
      </div>
      <div>Party A and Party B have reached the following agreement on house leasing：</div>
      <div style="margin: 8px 0 4px 0">一、租赁期限</div>
      <div>
        1、The lease term is from ______ month ______ day of ______ year to ______ month ______ day of ______ year。
      </div>
      <div>
        2、If Party B needs to continue leasing after the lease term expires, Party A has the right to consider renewing
        the lease with Party B first.
      </div>
      <div style="margin: 8px 0 4px 0">二、房屋使用费用</div>
      <div>1、Party B is responsible for the water, electricity, gas, and property management fees during the lease term.
      </div>
      <div>
        2、Party B should pay a deposit of ______元 to Party A in advance, which is used for repair and maintenance of the
        house during the lease term.
      </div>
      <div>
        3、Party B should keep the house and its equipment facilities in good condition during the lease term, and should
        not damage them by any means.
        If Party B causes damage to the house or its equipment facilities, Party B should pay the repair or replacement
        fee.
      </div>
      <div style="margin: 8px 0 4px 0">3. Requirements for Housing Use</div>
      <div>1、Party B should not use the house for illegal activities.</div>
      <div>2、Party B should protect the house and its equipment facilities, and should not arbitrarily modify or
        dismantle them.</div>
      <div>
        3、Party B should not rent the house to others, and should not use the house for storing dangerous goods and
        explosive materials.
      </div>
      <div style="margin: 8px 0 4px 0">四、支付方式</div>
      <div>1、Party B should pay the rent each month, and the rent is ______元/月。</div>
      <div>
        2、Party B should pay the rent to Party A on or before the ______ day of each month, and the payment method is
        (cash, bank transfer, etc.).
      </div>
      <div style="margin: 8px 0 4px 0">5. Terms of Contract Termination</div>
      <div>1、In the lease term, if Party B meets any of the following conditions, Party A has the right to terminate the
        contract:</div>
      <div>（1）Party B rents the house to others;</div>
      <div>（2）Party B uses the house for illegal activities;</div>
      <div>（3）Party B fails to pay the rent for 30 days or more.</div>
      <div>2、In the lease term, if Party A meets any of the following conditions, Party B has the right to terminate the
        contract:</div>
      <div>（1）Party A rents the house to others;</div>
      <div>（2）Party A uses the house for illegal activities;</div>
      <div>（3）Party A fails to provide the house and its equipment facilities for normal use for 30 days or more.</div>
      <div>3、If Party B does not renew the lease within 3 months after the end of the lease term, Party A has the right
        to return the house within 3 days.</div>
      <div>4、If Party B needs to continue leasing, Party B should apply for a new lease contract within 3 months of the
        end of the lease term, and the two parties should sign a new lease contract.</div>
      <div style="margin: 8px 0 4px 0">6. Other Matters</div>
      <div>1、This contract shall come into effect upon the signing and seal of both parties.</div>
      <div>2、This contract is valid for two copies, one for each party, and has the same legal effect.</div>
      <div style="margin-top: 28px">
        <span>Party A：</span>
        <span class="demo-input" style="min-width: 220px">
          {{ data.partyA }}
        </span>
        <span>&emsp;&emsp;Party B：</span>
        <span class="demo-input" style="min-width: 220px">
          {{ data.partyB }}
        </span>
      </div>
      <div>
        <span>Phone：</span>
        <span class="demo-input" style="min-width: 220px">{{ data.phone }}</span>
        <span>&emsp;&emsp;Phone：</span>
        <span class="demo-input" style="min-width: 220px">{{ data.phone }}</span>
      </div>
      <div>
        <span>Date：</span>
        <span class="demo-input" style="min-width: 220px">{{ data.date }}</span>
        <span>&emsp;&emsp;Date：</span>
        <span class="demo-input" style="min-width: 220px">{{ data.date }}</span>
      </div>
    </div>
    <div class="demo-title">QR code printing example: </div>
    <div style="text-align: center; page-break-after: always; padding: 40px 0">
      <ele-qr-code-svg v-if="data" :value="data.partyA" :size="180" />
      <div v-if="data" style="font-size: 14px; line-height: 1.2">
        {{ data.partyA }}
      </div>
    </div>
    <div class="demo-title">Table printing example：</div>
    <ele-table v-if="data && data.projects" size="large" :border="true" style="page-break-after: always">
      <colgroup>
        <col width="60px" />
        <col width="200px" />
        <col width="120px" />
        <col width="160px" />
        <col width="160px" />
        <col width="100px" />
      </colgroup>
      <thead>
        <tr>
          <th :rowspan="2"></th>
          <th :rowspan="2">Project Name</th>
          <th :rowspan="2">Amount</th>
          <th :colspan="2" style="text-align: center">Project Cycle</th>
          <th :rowspan="2">Status</th>
        </tr>
        <tr>
          <th>Start Time</th>
          <th>End Time</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(row, index) in data.projects" :key="row.id">
          <td style="text-align: center">{{ index + 1 }}</td>
          <td>{{ row.projectName }}</td>
          <td>{{ row.money }}(万元)</td>
          <td>{{ row.startDate }}</td>
          <td>{{ row.endDate }}</td>
          <td>
            <ele-text v-if="row.status === 0" type="success">In progress</ele-text>
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
        </tr>
      </tbody>
    </ele-table>
    <div class="demo-title">Image printing example：</div>
    <img src="https://cdn.eleadmin.com/20200610/4Z0QR2L0J1XStxBh99jVJ8qLfsGsOgjU.jpg"
      style="width: 100%; display: block; page-break-after: always" />
    <div class="demo-title">Table color modification：</div>
    <ele-table v-if="data && data.projects" size="large" :border="true" :style="{
      pageBreakAfter: 'always',
      '--ele-table-bg': '#fff',
      '--ele-table-color': '#000',
      '--ele-table-border-color': '#000',
      '--ele-table-th-color': '#000',
      '--ele-table-th-font-weight': 'bold',
      '--ele-table-th-bg': 'transparent',
      '--ele-table-tr-bg': 'transparent',
      '--ele-table-tr-hover-bg': 'transparent',
      '--ele-table-radius': 0,
      '--ele-table-sm-radius': 0,
      '--ele-table-lg-radius': 0,
      borderCollapse: 'collapse'
    }">
      <colgroup>
        <col width="60px" />
        <col width="200px" />
        <col width="120px" />
        <col width="160px" />
        <col width="160px" />
        <col width="100px" />
      </colgroup>
      <thead>
        <tr>
          <th :rowspan="2"></th>
          <th :rowspan="2">Project Name</th>
          <th :rowspan="2">Amount</th>
          <th :colspan="2" style="text-align: center">Project Cycle</th>
          <th :rowspan="2">Status</th>
        </tr>
        <tr>
          <th>Start Time</th>
          <th>End Time</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(row, index) in data.projects" :key="row.id">
          <td style="text-align: center">{{ index + 1 }}</td>
          <td>{{ row.projectName }}</td>
          <td>{{ row.money }}(万元)</td>
          <td>{{ row.startDate }}</td>
          <td>{{ row.endDate }}</td>
          <td>
            <span v-if="row.status === 0">In progress</span>
            <span v-else-if="row.status === 1">Delayed</span>
            <span v-else-if="row.status === 2">Not started</span>
            <span v-else-if="row.status === 3">Ended</span>
          </td>
        </tr>
      </tbody>
    </ele-table>
    <div class="demo-title">Watermark component：</div>
    <ele-watermark content="Ele Admin Plus" :gap="[60, 40]">
      <ele-table v-if="data && data.projects">
        <colgroup>
          <col width="38px" />
          <col width="200px" />
          <col width="140px" />
          <col width="138px" />
          <col width="138px" />
          <col width="118px" />
        </colgroup>
        <thead>
          <tr>
            <th></th>
            <th>Project Name</th>
            <th>Amount(万元)</th>
            <th>Start Time</th>
            <th>End Time</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(row, index) in data.projects" :key="row.id">
            <td style="text-align: center">{{ index + 1 }}</td>
            <td>{{ row.projectName }}</td>
            <td>{{ row.money }}(万元)</td>
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
          </tr>
        </tbody>
      </ele-table>
    </ele-watermark>
  </ele-printer>
</template>

<script setup>
import { ref } from 'vue';

defineProps({
  /** Contract Data */
  data: Object,
  /** Whether to always display */
  isStatic: Boolean,
  /** Print position */
  target: String
});

/** Printer component */
const printerRef = ref(null);

/** Print */
const print = () => {
  if (printerRef.value) {
    printerRef.value.print();
  }
};

defineExpose({ print });
</script>

<style lang="scss" scoped>
.demo-input {
  height: 24px;
  line-height: 24px;
  max-width: 80%;
  padding: 0 14px;
  display: inline-block;
  border-bottom: 1px solid #000;
  vertical-align: middle;
  box-sizing: border-box;
  overflow: hidden;
}

.demo-title {
  font-size: 22px;
  margin-bottom: 24px;
  line-height: 1.2;
}
</style>
