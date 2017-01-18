import {Config} from "../../Proxy/Proxy";

export function hasPermissionToShow(code: string, config: Config, isAdmin: boolean): boolean {
  if (isAdmin) {
    return true;
  }
  if (!config) {
    return false;
  }

  for (let permission of config.permissions) {
    if (permission.code === code) {
      return true;
    }
  }
  return false;
}
