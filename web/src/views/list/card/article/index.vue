<template>
  <ele-page>
    <top-search />
    <ele-card :body-style="{ padding: '16px 8px' }">
      <div class="article-list">
        <div v-for="item in data" :key="item.id" class="list-item">
          <div class="list-item-body">
            <div style="flex: 1">
              <ele-text size="md">{{ item.title }}</ele-text>
              <div style="margin-top: 8px; display: flex; align-items: center">
                <el-tag v-for="tag in item.tags" :key="tag" type="info" size="small" :disable-transitions="true"
                  style="margin-right: 8px">
                  {{ tag }}
                </el-tag>
              </div>
              <div style="margin-top: 16px">{{ item.content }}</div>
              <div style="margin-top: 16px; display: flex; align-items: center">
                <el-avatar :src="item.avatar" :size="24" />
                <ele-text type="secondary" style="flex: 1; padding-left: 6px">
                  {{ item.user }} published {{ item.time }}
                </ele-text>
              </div>
              <div style="margin-top: 16px; display: flex; align-items: center">
                <ele-text :icon="Star" type="placeholder" style="cursor: pointer">
                  &nbsp;{{ item.likes }}
                </ele-text>
                <el-divider direction="vertical" style="margin: 0 14px" />
                <ele-text :icon="CollectionTag" type="placeholder" style="cursor: pointer">
                  &nbsp;{{ item.favorites }}
                </ele-text>
                <el-divider direction="vertical" style="margin: 0 14px" />
                <ele-text :icon="ChatLineSquare" type="placeholder" style="cursor: pointer">
                  &nbsp;{{ item.comments }}
                </ele-text>
              </div>
            </div>
            <div class="list-item-extra">
              <img-grid :images="item.covers" />
            </div>
          </div>
          <el-divider style="margin: 0; opacity: 0.6" />
        </div>
      </div>
      <div style="margin-top: 16px; text-align: center">
        <el-button v-if="page !== 1" :loading="loading" @click="query">
          {{ loading ? 'Loading..' : 'Load more' }}
        </el-button>
      </div>
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref } from 'vue';
import { Star, CollectionTag, ChatLineSquare } from '@element-plus/icons-vue';
import TopSearch from '../project/components/top-search.vue';
import ImgGrid from './components/img-grid.vue';

const data = ref([]);

const loading = ref(false);

const page = ref(2);

const query = () => {
  loading.value = true;
  setTimeout(() => {
    loading.value = false;
    page.value++;
    data.value = data.value.concat(
      data.value.slice(0, 3).map((d, i) => {
        return { ...d, id: data.value.length + i + 1 };
      })
    );
  }, 1000);
};

data.value = [
  {
    id: 1,
    title: 'ElementUI',
    content:
      'Element, a Vue 2.0 component library for developers, designers, and product managers, provides supporting design resources to help you quickly build your website.。',
    time: '2 小时前',
    covers: [
      'https://cdn.eleadmin.com/20200610/RZ8FQmZfHkcffMlTBCJllBFjEhEsObVo.jpg',
      'https://cdn.eleadmin.com/20200610/WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg',
      'https://cdn.eleadmin.com/20200610/4Z0QR2L0J1XStxBh99jVJ8qLfsGsOgjU.jpg',
      'https://cdn.eleadmin.com/20200610/ttkIjNPlVDuv4lUTvRX8GIlM2QqSe0jg.jpg',
      'https://cdn.eleadmin.com/20200610/fAenQ8nvRjL7x0i0jEfuDBZHvJfHf3v6.jpg',
      'https://cdn.eleadmin.com/20200610/LrCTN2j94lo9N7wEql7cBr1Ux4rHMvmZ.jpg'
    ],
    tags: ['EleAdminPro', 'UI framework', 'design language'],
    user: 'SunSmile',
    avatar:
      'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg',
    favorites: 104,
    likes: 189,
    comments: 15
  },
  {
    id: 2,
    title: 'Vue.js',
    content:
      'Vue, a progressive framework for building user interfaces, is designed to be incrementally adoptable. ',
    time: '4 小时前',
    covers: [
      'https://cdn.eleadmin.com/20200610/WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg',
      'https://cdn.eleadmin.com/20200610/4Z0QR2L0J1XStxBh99jVJ8qLfsGsOgjU.jpg',
      'https://cdn.eleadmin.com/20200610/ttkIjNPlVDuv4lUTvRX8GIlM2QqSe0jg.jpg',
      'https://cdn.eleadmin.com/20200610/fAenQ8nvRjL7x0i0jEfuDBZHvJfHf3v6.jpg',
      'https://cdn.eleadmin.com/20200610/LrCTN2j94lo9N7wEql7cBr1Ux4rHMvmZ.jpg',
      'https://cdn.eleadmin.com/20200610/yeKvhT20lMU0f1T3Y743UlGEOLLnZSnp.jpg'
    ],
    tags: ['Vue', 'UI framework', 'design language'],
    user: 'Vue.js',
    avatar:
      'https://cdn.eleadmin.com/20200609/b6a811873e704db49db994053a5019b2.jpg',
    favorites: 104,
    likes: 189,
    comments: 15
  },
  {
    id: 3,
    title: 'Vuex',
    content:
      'Vuex is a state management pattern developed specifically for Vue.js applications. It uses a centralized storage to manage the state of all components of the application and uses corresponding rules to ensure that the state changes in a predictable way.。',
    time: '12 小时前',
    covers: [
      'https://cdn.eleadmin.com/20200610/4Z0QR2L0J1XStxBh99jVJ8qLfsGsOgjU.jpg',
      'https://cdn.eleadmin.com/20200610/ttkIjNPlVDuv4lUTvRX8GIlM2QqSe0jg.jpg',
      'https://cdn.eleadmin.com/20200610/fAenQ8nvRjL7x0i0jEfuDBZHvJfHf3v6.jpg',
      'https://cdn.eleadmin.com/20200610/LrCTN2j94lo9N7wEql7cBr1Ux4rHMvmZ.jpg',
      'https://cdn.eleadmin.com/20200610/yeKvhT20lMU0f1T3Y743UlGEOLLnZSnp.jpg',
      'https://cdn.eleadmin.com/20200610/CyrCNmTJfv7D6GFAg39bjT3eRkkRm5dI.jpg'
    ],
    tags: ['Vue', 'UI framework', 'design language'],
    user: 'Vue.js',
    avatar:
      'https://cdn.eleadmin.com/20200609/948344a2a77c47a7a7b332fe12ff749a.jpg',
    favorites: 104,
    likes: 189,
    comments: 15
  },
  {
    id: 4,
    title: 'Vue Router',
    content:
      'Vue Router is the official routing manager for Vue.js. It is deeply integrated with the core of Vue.js, making building single- page applications a breeze.',
    time: '14 小时前',
    covers: [
      'https://cdn.eleadmin.com/20200610/ttkIjNPlVDuv4lUTvRX8GIlM2QqSe0jg.jpg',
      'https://cdn.eleadmin.com/20200610/fAenQ8nvRjL7x0i0jEfuDBZHvJfHf3v6.jpg',
      'https://cdn.eleadmin.com/20200610/LrCTN2j94lo9N7wEql7cBr1Ux4rHMvmZ.jpg',
      'https://cdn.eleadmin.com/20200610/yeKvhT20lMU0f1T3Y743UlGEOLLnZSnp.jpg',
      'https://cdn.eleadmin.com/20200610/CyrCNmTJfv7D6GFAg39bjT3eRkkRm5dI.jpg',
      'https://cdn.eleadmin.com/20200610/RZ8FQmZfHkcffMlTBCJllBFjEhEsObVo.jpg'
    ],
    tags: ['Vue', 'UI framework', 'design language'],
    user: 'Jasmine',
    avatar:
      'https://cdn.eleadmin.com/20200609/f6bc05af944a4f738b54128717952107.jpg',
    favorites: 104,
    likes: 189,
    comments: 15
  },
  {
    id: 5,
    title: 'Sass',
    content: 'SSass is the world is most mature, stable, and powerful professional- grade CSS extension language',
    time: '10 hours ago',
    covers: [
      'https://cdn.eleadmin.com/20200610/fAenQ8nvRjL7x0i0jEfuDBZHvJfHf3v6.jpg',
      'https://cdn.eleadmin.com/20200610/LrCTN2j94lo9N7wEql7cBr1Ux4rHMvmZ.jpg',
      'https://cdn.eleadmin.com/20200610/yeKvhT20lMU0f1T3Y743UlGEOLLnZSnp.jpg',
      'https://cdn.eleadmin.com/20200610/CyrCNmTJfv7D6GFAg39bjT3eRkkRm5dI.jpg',
      'https://cdn.eleadmin.com/20200610/RZ8FQmZfHkcffMlTBCJllBFjEhEsObVo.jpg',
      'https://cdn.eleadmin.com/20200610/WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg'
    ],
    tags: ['Vue', 'UI framework', 'design language'],
    user: 'Coolest',
    avatar:
      'https://cdn.eleadmin.com/20200609/2d98970a51b34b6b859339c96b240dcd.jpg',
    favorites: 104,
    likes: 189,
    comments: 15
  },
  {
    id: 6,
    title: 'Axios',
    content:
      'Axios is a promise-based HTTP library that can be used in browsers and node.js.。',
    time: '16 hours ago',
    covers: [
      'https://cdn.eleadmin.com/20200610/LrCTN2j94lo9N7wEql7cBr1Ux4rHMvmZ.jpg',
      'https://cdn.eleadmin.com/20200610/yeKvhT20lMU0f1T3Y743UlGEOLLnZSnp.jpg',
      'https://cdn.eleadmin.com/20200610/CyrCNmTJfv7D6GFAg39bjT3eRkkRm5dI.jpg',
      'https://cdn.eleadmin.com/20200610/RZ8FQmZfHkcffMlTBCJllBFjEhEsObVo.jpg',
      'https://cdn.eleadmin.com/20200610/WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg',
      'https://cdn.eleadmin.com/20200610/4Z0QR2L0J1XStxBh99jVJ8qLfsGsOgjU.jpg'
    ],
    tags: ['Vue', 'UI framework', 'design language'],
    user: 'SunSmile',
    avatar:
      'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg',
    favorites: 104,
    likes: 189,
    comments: 15
  },
  {
    id: 7,
    title: 'Webpack',
    content:
      'webpack is a module bundler. The main goal of webpack is to bundle JavaScript files together, and the bundled files are used in the browser.',
    time: '6 hours ago',
    covers: [
      'https://cdn.eleadmin.com/20200610/yeKvhT20lMU0f1T3Y743UlGEOLLnZSnp.jpg',
      'https://cdn.eleadmin.com/20200610/CyrCNmTJfv7D6GFAg39bjT3eRkkRm5dI.jpg',
      'https://cdn.eleadmin.com/20200610/RZ8FQmZfHkcffMlTBCJllBFjEhEsObVo.jpg',
      'https://cdn.eleadmin.com/20200610/WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg',
      'https://cdn.eleadmin.com/20200610/4Z0QR2L0J1XStxBh99jVJ8qLfsGsOgjU.jpg',
      'https://cdn.eleadmin.com/20200610/ttkIjNPlVDuv4lUTvRX8GIlM2QqSe0jg.jpg'
    ],
    tags: ['Vue', 'UI framework', 'design language'],
    user: '全村人的希望',
    avatar:
      'https://cdn.eleadmin.com/20200609/948344a2a77c47a7a7b332fe12ff749a.jpg',
    favorites: 104,
    likes: 189,
    comments: 15
  },
  {
    id: 8,
    title: 'Node.js',
    content:
      'Node.js is a JavaScript runtime environment based on the Chrome V8 engine. Node.js uses an event-driven, non-blocking I/O model, making it lightweight and efficient.',
    time: '8 hours ago',
    covers: [
      'https://cdn.eleadmin.com/20200610/CyrCNmTJfv7D6GFAg39bjT3eRkkRm5dI.jpg',
      'https://cdn.eleadmin.com/20200610/RZ8FQmZfHkcffMlTBCJllBFjEhEsObVo.jpg',
      'https://cdn.eleadmin.com/20200610/WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg',
      'https://cdn.eleadmin.com/20200610/4Z0QR2L0J1XStxBh99jVJ8qLfsGsOgjU.jpg',
      'https://cdn.eleadmin.com/20200610/ttkIjNPlVDuv4lUTvRX8GIlM2QqSe0jg.jpg',
      'https://cdn.eleadmin.com/20200610/fAenQ8nvRjL7x0i0jEfuDBZHvJfHf3v6.jpg'
    ],
    tags: ['Vue', 'UI framework', 'design language'],
    user: 'Jasmine',
    avatar:
      'https://cdn.eleadmin.com/20200609/f6bc05af944a4f738b54128717952107.jpg',
    favorites: 104,
    likes: 189,
    comments: 15
  }
];
</script>

<script>
export default {
  name: 'ListCardArticle'
};
</script>

<style lang="scss" scoped>
.list-item-body {
  display: flex;
  padding: 16px 24px;
}

.list-item-extra {
  width: 280px;
  height: 175px;
  margin-left: 12px;
}

@media screen and (max-width: 880px) {
  .list-item-extra {
    width: 200px;
    height: 125px;
  }
}

@media screen and (max-width: 576px) {
  .list-item-body {
    display: block;
  }

  .list-item-extra {
    width: 100%;
    height: 200px;
    margin: 16px 0 0 0;
  }
}
</style>
