import { redirect } from "@sveltejs/kit";
import { userStore } from "$lib/stores/user";

export function load() {
  userStore.set(null);

  throw redirect(307, "/");
}

