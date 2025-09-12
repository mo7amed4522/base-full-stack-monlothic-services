import request from '@/utils/request';

/**
 * Query unread notifications
 */
export async function getUnreadNotice() {
  const res = await request.get('https://cdn.eleadmin.com/20200610/message.json');
  return res.data;
}

/**
 * Pagination query notification
 */
export async function pageNotices(_params) {
  const res = await request.get('https://cdn.eleadmin.com/20200610/notices.json');
  return res.data;
}

/**
 * Pagination query private letter
 */
export async function pageLetters(_params) {
  const res = await request.get('https://cdn.eleadmin.com/20200610/letters.json');
  return res.data;
}

/**
 * Pagination query todo
 */
export async function pageTodos(_params) {
  const res = await request.get('https://cdn.eleadmin.com/20200610/todos.json');
  return res.data;
}

/**
 * Query unread number
 */
export async function getUnReadNum() {
  return { notice: 2, letter: 3, todo: 4 };
}

/**
 * Get added members
 */
export async function listAddedUsers() {
  const data = [
    {
      key: '1',
      number: '00001',
      name: 'John Brown',
      department: 'R&D Department'
    },
    {
      key: '2',
      number: '00002',
      name: 'Jim Green',
      department: 'Product Department'
    },
    {
      key: '3',
      number: '00003',
      name: 'Joe Black',
      department: 'Product Department'
    }
  ];
  return data;
}

/**
 * Get user scores
 */
export async function pageUserScores() {
  const res = await request.get('https://cdn.eleadmin.com/20200610/example-table-merge.json');
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Get volume list
 */
export async function getPieceList(params) {
  const res = await request.get('https://cdn.eleadmin.com/20200610/document.json', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Query project progress
 */
export async function getProjectList() {
  const res = await request.get('https://cdn.eleadmin.com/20200610/project-list.json');
  return res.data;
}

/**
 * Get demo markdown
 */
export async function getDemoMD() {
  const res = await request.get('https://cdn.eleadmin.com/20200610/demo.md');
  return res.data;
}

/**
 * Get demo multiple selection card
 */
export async function getDemoCardData() {
  return [
    {
      value: 1,
      title: 'ElementUI',
      content: 'EElement is a Vue 2.0 component library for developers, designers, and product managers. It provides supporting design resources to help you quickly build your website.',
      avatar: 'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg'
    },
    {
      value: 2,
      title: 'Vue.js',
      content: 'Vue is a progressive framework for building user interfaces. Unlike other large frameworks, Vue is designed to be applied layer by layer from the bottom up.',
      avatar: 'https://cdn.eleadmin.com/20200609/b6a811873e704db49db994053a5019b2.jpg'
    },
    {
      value: 3,
      title: 'Vue Router',
      content: 'Vue Router is the official routing manager for Vue.js. It is deeply integrated with the core of Vue.js, making building single-page applications a breeze.',
      avatar: 'https://cdn.eleadmin.com/20200609/f6bc05af944a4f738b54128717952107.jpg'
    }
  ];
}
