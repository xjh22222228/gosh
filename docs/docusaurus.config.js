module.exports = {
  title: 'GOSH',
  tagline: '👌 Golang utility library, With additional functions such as JavaScript/Python!',
  url: 'https://github.com/xjh22222228/gosh',
  baseUrl: '/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.svg',
  organizationName: 'xjh22222228', // Usually your GitHub org/user name.
  projectName: 'gosh', // Usually your repo name.
  themeConfig: {
    navbar: {
      logo: {
        alt: 'GOSH Logo',
        src: 'img/logo.svg',
      },
      items: [
        {
          to: 'docs/',
          activeBasePath: 'docs',
          label: 'Documentation',
          position: 'left',
        },
        {
          href: 'https://github.com/xjh22222228/gosh',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Resources',
          items: [
            {
              label: 'Nav',
              href: 'https://github.com/xjh22222228/nav',
            },
            {
              label: 'Tomato Work',
              href: 'https://github.com/xjh22222228/tomato-work',
            },
            {
              label: 'Bo💣mb',
              href: 'https://github.com/xjh22222228/boomb',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Golang',
              href: 'https://golang.org/',
            },
            {
              label: 'GOSH',
              href: 'https://github.com/xjh22222228/gosh',
            },
            {
              label: 'Docusaurus',
              href: 'https://v2.docusaurus.io/',
            }
          ],
        },
        {
          title: 'Support',
          items: [
            {
              label: 'Author',
              href: 'https://github.com/xjh22222228',
            },
          ],
        }
      ],
      copyright: `Copyright © 2021-${new Date().getFullYear()} GOSH, Inc. Built with Docusaurus.`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl:
            'https://github.com/xjh22222228/gosh/edit/docs/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
