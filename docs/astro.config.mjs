import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";
import { rehypeMermaid } from "@beoe/rehype-mermaid";
import { getCache } from "@beoe/cache";

const cache = getCache();

export default defineConfig({
  site: "https://weburz.github.io/crisp",
  base: "crisp",
  integrations: [
    starlight({
      title: "Crisp",
      description: "A linter for Git commit messages.",
      editLink: {
        baseUrl: "https://github.com/Weburz/crisp/edit/main/docs",
      },
      social: {
        github: "https://github.com/Weburz/crisp",
        discord: "https://discord.gg/QeYqwyxBhR",
        email: "mailto:contact@weburz.com",
        facebook: "https://www.facebook.com/Weburz",
        instagram: "https://www.instagram.com/weburzit",
        linkedin: "https://www.linkedin.com/company/weburz",
        youtube: "https://www.youtube.com/@Weburz",
        twitter: "https://x.com/weburz",
      },
      lastUpdated: true,
      sidebar: [
        {
          label: "Welcome",
          autogenerate: {
            directory: "usage-guide",
          },
        },
        {
          label: "Developer Guide",
          autogenerate: {
            directory: "dev-guide",
          },
        },
        {
          label: "Software Architecture",
          autogenerate: {
            directory: "architecture",
          },
        },
      ],
      credits: true,
      components: {
        // TODO: Identify how to get it to work else it breaks compilation
        // Head: "./src/components/Head.astro",
      },
    }),
  ],
  markdown: {
    rehypePlugins: [
      [
        rehypeMermaid,
        {
          strategy: "file",
          fsPath: "public/beoe",
          webPath: "/crisp/beoe",
          cache,
        },
      ],
    ],
  },
});
