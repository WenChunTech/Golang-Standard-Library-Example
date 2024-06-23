import { defineConfig } from 'vitepress'
import { generateSidebar } from 'vitepress-sidebar';

const vitepressSidebarOptions = {
    /*
 * For detailed instructions, see the links below:
 * https://github.com/jooy2/vitepress-sidebar#options
 */
    documentRootPath: 'docs',
    // scanStartPath: null,
    // resolvePath: null,
    // useTitleFromFileHeading: true,
    // useTitleFromFrontmatter: true,
    // useFolderTitleFromIndexFile: false,
    // useFolderLinkFromIndexFile: false,
    // hyphenToSpace: true,
    // underscoreToSpace: true,
    // capitalizeFirst: false,
    // capitalizeEachWords: false,
    // collapsed: true,
    // collapseDepth: 2,
    // sortMenusByName: false,
    // sortMenusByFrontmatterOrder: false,
    // sortMenusByFrontmatterDate: false,
    // sortMenusOrderByDescending: false,
    // sortMenusOrderNumericallyFromTitle: false,
    // sortMenusOrderNumericallyFromLink: false,
    // frontmatterOrderDefaultValue: 0,
    // manualSortFileNameByPriority: ['first.md', 'second', 'third.md'],
    // removePrefixAfterOrdering: false,
    // prefixSeparator: '.',
    // excludeFiles: ['first.md', 'secret.md'],
    // excludeFilesByFrontmatterFieldName: 'exclude',
    // excludeFolders: ['secret-folder'],
    // includeDotFiles: false,
    // includeRootIndexFile: false,
    // includeFolderIndexFile: false,
    // includeEmptyFolder: false,
    // rootGroupText: 'Contents',
    // rootGroupLink: 'https://github.com/jooy2',
    // rootGroupCollapsed: false,
    // convertSameNameSubFileToGroupIndexPage: false,
    // folderLinkNotIncludesFileName: false,
    // keepMarkdownSyntaxFromTitle: false,
    // debugPrint: false,
};

// https://vitepress.dev/reference/site-config
export default defineConfig({
    title: "首页",
    description: "the Golang Standard Library Example",
    themeConfig: {
        // https://vitepress.dev/reference/default-theme-config
        nav: [
            { text: 'Home', link: '/' },
            { text: 'Examples', link: '/golang/markdown-examples' }
        ],

        sidebar: generateSidebar(vitepressSidebarOptions),

        socialLinks: [
            { icon: 'github', link: 'https://github.com/vuejs/vitepress' }
        ]
    }
})
