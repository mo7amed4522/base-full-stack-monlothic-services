import { ref } from 'vue';
import ExcelJS from 'exceljs';
import { download } from '@/utils/common';

/**
 * Component global configuration
 */
export function useGlobalConfig() {
  /** Global configuration of advanced tables */
  const tableConfig = ref({
    response: {
      dataName: 'list',
      countName: 'total'
    },
    // size: ref('small'),
    tools: ['reload', 'export', 'size', 'columns', 'maximized'],
    exportConfig: {
      // Use exceljs for export
      beforeExport: (params) => {
        const { fileName, closeModal } = params;
        const workbook = getExportWorkbook(params);
        workbook.xlsx.writeBuffer().then((data) => {
          download(data, `${fileName}.xlsx`);
          closeModal();
        });
        return false;
      }
    }
  });

  return { tableConfig };
}

/**
 * Generate exceljs workbook for export
 * @param params Parameters
 */
export function getExportWorkbook(params) {
  const { headerData, bodyData, footerData, bodyCols } = params;
  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet('Sheet1');
  // Traversing the data
  const sheetRows = [];
  const sheetStyles = [];
  const sheetMerges = [];
  const sheetRowHeight = [];
  [...headerData, ...bodyData, ...footerData].forEach((item, index) => {
    const sheetCols = [];
    const rowStyles = [];
    item.forEach((d, colIndex) => {
      if (!d.isTreeCell || !d.indent) {
        sheetCols.push(d.text ?? '');
      } else {
        // Tree table indentation
        const gap = Array.from({ length: d.indent }).fill('  ').join('');
        sheetCols.push(gap + (d.text ?? ''));
      }
      // Cell style
      rowStyles.push({
        font: { size: 12, bold: index < headerData.length },
        alignment: {
          vertical: d.isTreeIndex && d.rowspan !== 1 ? 'top' : 'middle',
          horizontal: d.isExpandCell ? 'left' : d.column?.align || 'left',
          wrapText: d.isExpandCell ? true : false,
          indent: d.isExpandCell ? 1 : 0
        },
        border: {
          top: { style: 'thin' },
          left: d.hideLeftBorder ? {} : { style: 'thin' },
          bottom: { style: 'thin' },
          right: d.hideRightBorder ? {} : { style: 'thin' }
        }
      });
      // Cell merge
      if ((d.rowspan && d.rowspan > 1) || (d.colspan && d.colspan > 1)) {
        sheetMerges.push([index + 1, colIndex + 1, index + (d.rowspan || 1), colIndex + (d.colspan || 1)]);
      }
    });
    sheetRows.push(sheetCols);
    sheetStyles.push(rowStyles);
    // Row height
    const expandCell = item.find((d) => d.isExpandCell && d.colspan);
    const lines = expandCell?.text?.split?.('\n')?.length || 1;
    sheetRowHeight.push(lines * 20);
  });
  // Add to the table
  sheet.addRows(sheetRows).forEach((row, index) => {
    row.eachCell({ includeEmpty: true }, (cell, colIndex) => {
      Object.assign(cell, sheetStyles[index][colIndex - 1]);
    });
    row.height = sheetRowHeight[index];
  });
  sheetMerges.forEach((merges) => {
    sheet.mergeCells(merges);
  });
  // Column width
  bodyCols.forEach((col, colIndex) => {
    const w = col.width ?? col.minWidth;
    sheet.getColumn(colIndex + 1).width = w == null ? void 0 : w / 8;
  });
  return workbook;
}
