import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'go-cube',
  description: 'CubeJS 兼容的 Go 实现，后端为 ClickHouse',
  base: '/go-cube/',
  themeConfig: {
    nav: [
      { text: '首页', link: '/' },
      { text: '设计文档', link: '/design/sql-builder' },
    ],
    sidebar: [
      {
        text: '设计文档',
        items: [
          { text: '前端设计', link: '/design/design' },
          { text: 'SQL 拼接设计', link: '/design/sql-builder' },
        ],
      },
    ],
    socialLinks: [
      { icon: 'github', link: 'https://github.com/Servicewall/go-cube' },
    ],
    editLink: {
      pattern: 'https://github.com/Servicewall/go-cube/edit/main/docs/:path',
      text: '在 GitHub 上编辑此页',
    },
  },
})
