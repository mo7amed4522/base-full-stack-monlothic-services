<template>
  <ele-page :multi-card="false">
    <ele-card>
      <el-result icon="success" title="提交成功">
        <template #sub-title>
          <ele-text type="placeholder">
            The submission result page is used to feed back the processing results of a series of operation tasks. The
            gray area can display some supplementary information.
          </ele-text>
        </template>
      </el-result>
      <div style="text-align: center">
        <el-button type="primary">Return to list</el-button>
        <el-button>View project</el-button>
        <el-button>Print</el-button>
      </div>
      <div style="
          max-width: 740px;
          padding: 20px 32px;
          background: hsla(0, 0%, 60%, 0.1);
          border-radius: var(--el-border-radius-base);
          margin: 26px auto;
        ">
        The application has been submitted and is waiting for department approval.
      </div>
    </ele-card>
  </ele-page>
</template>

<script>
export default {
  name: 'ResultSuccess'
};
</script>
