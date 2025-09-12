<template>
  <ele-page :multi-card="false">
    <top-search />
    <el-row :gutter="16">
      <el-col v-for="item in data" :key="item.id" :lg="6" :md="8" :sm="12" :xs="24">
        <ele-card shadow="hover" :body-style="{ padding: 0 }"
          style="margin-top: 16px; overflow: hidden; cursor: pointer">
          <img :src="item.cover" style="width: 100%" />
          <div style="padding: 14px 24px 24px 24px">
            <ele-text size="md" style="max-height: 100%">
              {{ item.title }}
            </ele-text>
            <ele-ellipsis :max-line="2" type="placeholder" style="margin: 8px 0 16px 0; max-height: 100%">
              {{ item.content }}
            </ele-ellipsis>
            <div style="display: flex; align-items: center">
              <ele-text type="placeholder" style="flex: 1">
                {{ item.time }}
              </ele-text>
              <ele-avatar-group size="small" :data="item.users" style="flex-shrink: 0" />
            </div>
          </div>
        </ele-card>
      </el-col>
    </el-row>
    <ele-pagination :total="count" v-model:page-size="limit" v-model:current-page="page" layout="prev, pager, next"
      prev-text="Previous page" next-text="Next page" :style="{
        marginTop: '18px',
        justifyContent: 'center',
        '--ele-pagination-hover-bg': 'transparent',
        '--ele-pagination-hover-color': 'var(--el-color-primary)'
      }" />
  </ele-page>
</template>

<script setup>
import { ref } from 'vue';
import TopSearch from './components/top-search.vue';

/** 数据 */
const data = ref([]);

/** 第几页 */
const page = ref(1);

/** 每页多少条 */
const limit = ref(8);

/** 总数量 */
const count = ref(0);

/** 查询数据 */
const query = () => {
  count.value = 40;
  data.value = [
    {
      id: 1,
      title: 'ElementUI',
      content:
        'Element is a component library for Vue.js. It is a set of components that help you build websites faster and easier.',
      time: '2 hours ago',
      cover:
        'https://cdn.eleadmin.com/20200610/RZ8FQmZfHkcffMlTBCJllBFjEhEsObVo.jpg',
      users: [
        {
          label: 'SunSmile',
          value:
            'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg'
        },
        {
          label: '酷酷的大叔',
          value:
            'https://cdn.eleadmin.com/20200609/2d98970a51b34b6b859339c96b240dcd.jpg'
        },
        {
          label: 'Jasmine',
          value:
            'https://cdn.eleadmin.com/20200609/948344a2a77c47a7a7b332fe12ff749a.jpg'
        }
      ]
    },
    {
      id: 2,
      title: 'Vue.js',
      content:
        'Vue is a progressive framework for building user interfaces. Unlike other monolithic frameworks, Vue is designed from the ground up to be incrementally adoptable.',
      time: '4 hours ago',
      cover:
        'https://cdn.eleadmin.com/20200610/WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg',
      users: [
        {
          label: 'SunSmile',
          value:
            'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg'
        },
        {
          label: 'Coolest',
          value:
            'https://cdn.eleadmin.com/20200609/2d98970a51b34b6b859339c96b240dcd.jpg'
        },
        {
          label: 'Jasmine',
          value:
            'https://cdn.eleadmin.com/20200609/948344a2a77c47a7a7b332fe12ff749a.jpg'
        }
      ]
    },
    {
      id: 3,
      title: 'Vuex',
      content:
        'Vuex is a state management pattern + library for Vue.js applications. It serves as a centralized store for all the components in an application, with rules ensuring that the state can only be mutated in a predictable fashion.',
      time: '12 hours ago',
      cover:
        'https://cdn.eleadmin.com/20200610/4Z0QR2L0J1XStxBh99jVJ8qLfsGsOgjU.jpg',
      users: [
        {
          label: 'SunSmile',
          value:
            'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg'
        },
        {
          label: 'Coolest',
          value:
            'https://cdn.eleadmin.com/20200609/2d98970a51b34b6b859339c96b240dcd.jpg'
        },
        {
          label: 'Jasmine',
          value:
            'https://cdn.eleadmin.com/20200609/948344a2a77c47a7a7b332fe12ff749a.jpg'
        }
      ]
    },
    {
      id: 4,
      title: 'Vue Router',
      content:
        'Vue Router is the official router for Vue.js. It deeply integrates with Vue.js core to make building Single Page Applications with Vue.js a breeze.',
      time: '14 hours ago',
      cover:
        'https://cdn.eleadmin.com/20200610/ttkIjNPlVDuv4lUTvRX8GIlM2QqSe0jg.jpg',
      users: [
        {
          label: 'SunSmile',
          value:
            'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg'
        },
        {
          label: 'Coolest',
          value:
            'https://cdn.eleadmin.com/20200609/2d98970a51b34b6b859339c96b240dcd.jpg'
        },
        {
          label: 'Jasmine',
          value:
            'https://cdn.eleadmin.com/20200609/948344a2a77c47a7a7b332fe12ff749a.jpg'
        }
      ]
    },
    {
      id: 5,
      title: 'Sass',
      content:
        'Sass is a pre-processor scripting language that is compiled into CSS. It is a powerful and elegant way to write CSS.',
      time: '10 hours ago',
      cover:
        'https://cdn.eleadmin.com/20200610/fAenQ8nvRjL7x0i0jEfuDBZHvJfHf3v6.jpg',
      users: [
        {
          label: 'SunSmile',
          value:
            'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg'
        },
        {
          label: 'Coolest',
          value:
            'https://cdn.eleadmin.com/20200609/2d98970a51b34b6b859339c96b240dcd.jpg'
        },
        {
          label: 'Jasmine',
          value:
            'https://cdn.eleadmin.com/20200609/948344a2a77c47a7a7b332fe12ff749a.jpg'
        }
      ]
    },
    {
      id: 6,
      title: 'Axios',
      content:
        'Axios is a promise-based HTTP client for the browser and node.js.',
      time: '16 hours ago',
      cover:
        'https://cdn.eleadmin.com/20200610/LrCTN2j94lo9N7wEql7cBr1Ux4rHMvmZ.jpg',
      users: [
        {
          label: 'SunSmile',
          value:
            'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg'
        },
        {
          label: 'Coolest',
          value:
            'https://cdn.eleadmin.com/20200609/2d98970a51b34b6b859339c96b240dcd.jpg'
        },
        {
          label: 'Jasmine',
          value:
            'https://cdn.eleadmin.com/20200609/948344a2a77c47a7a7b332fe12ff749a.jpg'
        }
      ]
    },
    {
      id: 7,
      title: 'Webpack',
      content:
        'Webpack is a module bundler. Webpack takes modules with dependencies and generates static assets representing those modules.',
      time: '6 hours ago',
      cover:
        'https://cdn.eleadmin.com/20200610/yeKvhT20lMU0f1T3Y743UlGEOLLnZSnp.jpg',
      users: [
        {
          label: 'SunSmile',
          value:
            'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg'
        },
        {
          label: 'Coolest',
          value:
            'https://cdn.eleadmin.com/20200609/2d98970a51b34b6b859339c96b240dcd.jpg'
        },
        {
          label: 'Jasmine',
          value:
            'https://cdn.eleadmin.com/20200609/948344a2a77c47a7a7b332fe12ff749a.jpg'
        }
      ]
    },
    {
      id: 8,
      title: 'Node.js',
      content:
        'Node.js is a JavaScript runtime built on Chrome\'s V8 JavaScript engine. Node.js uses an event-driven, non-blocking I/O model that makes it lightweight and efficient.',
      time: '8 hours ago',
      cover:
        'https://cdn.eleadmin.com/20200610/CyrCNmTJfv7D6GFAg39bjT3eRkkRm5dI.jpg',
      users: [
        {
          label: 'SunSmile',
          value:
            'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg'
        },
        {
          label: 'Coolest',
          value:
            'https://cdn.eleadmin.com/20200609/2d98970a51b34b6b859339c96b240dcd.jpg'
        },
        {
          label: 'Jasmine',
          value:
            'https://cdn.eleadmin.com/20200609/948344a2a77c47a7a7b332fe12ff749a.jpg'
        }
      ]
    }
  ];
};

// query();
</script>

<script>
export default {
  name: 'ListCardProject'
};
</script>
