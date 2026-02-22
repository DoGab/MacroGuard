import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ cookies }) => {
  // shadcn-svelte sidebar uses this default cookie name
  const sidebarCookie = cookies.get("sidebar:state");
  const defaultOpen = sidebarCookie ? sidebarCookie === "true" : true;

  return {
    sidebarOpen: defaultOpen
  };
};
