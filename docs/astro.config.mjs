import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";

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
      ],
    }),
  ],
});
