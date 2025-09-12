<template>
  <div style="padding: 8px 16px">
    <ele-table :border="true" :print-skin="true">
      <colgroup>
        <col width="50px" />
        <col />
        <col />
        <col width="110px" />
        <col />
      </colgroup>
      <thead>
        <tr>
          <th></th>
          <th>Data Name</th>
          <th>Data Code</th>
          <th style="text-align: center">Sort Number</th>
          <th style="text-align: center">Create Time</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(row, index) in data" :key="row.dictDataId">
          <td style="text-align: center">{{ index + 1 }}</td>
          <td>{{ row.dictDataName }}</td>
          <td>{{ row.dictDataCode }}</td>
          <td style="text-align: center">{{ row.sortNumber }}</td>
          <td style="text-align: center">{{ row.createTime }}</td>
        </tr>
      </tbody>
    </ele-table>
  </div>
</template>

<script setup>
defineProps({
  data: Array
});
</script>
