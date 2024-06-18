import { userStore } from "$lib/stores/user";
import { redirect } from "@sveltejs/kit";

export function load() {
  userStore.set({ username: "John Smith", accessToken: "foo" });
  throw redirect(307, "/me");
}
