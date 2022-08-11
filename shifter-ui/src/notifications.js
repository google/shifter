/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { useToast } from "vue-toastification";
// Get toast interface
const toast = useToast();

export function notifyAxiosError(error, msg = null, timeout = 4000) {
  var response = "Error";
  if (error.message !== undefined && error.message !== null) {
    response = error.message;
  }

  if (msg !== null) {
    response = msg + " - " + response;
  }
  toast.error(response, {
    timeout: timeout,
  });
}

export function shifterConversionSuccess(msg = null, timeout = 5000) {
  var response = "Success";
  if (msg !== null) {
    response = response + " - " + msg;
  }
  toast.success(response, {
    timeout: timeout,
  });
}

export function shifterConfigurationUpdateSuccess(msg = null, timeout = 2000) {
  var response = "Success";
  if (msg !== null) {
    response = response + " - " + msg;
  }
  toast.success(response, {
    timeout: timeout,
  });
}

export function shifterConfigurationUpdateError(msg = null, timeout = 2000) {
  var response = "Error";
  if (msg !== null) {
    response = msg + " - " + response;
  }
  toast.error(response, {
    timeout: timeout,
  });
}

// optionally export a default object
export default {
  shifterConversionSuccess,
  notifyAxiosError,
  shifterConfigurationUpdateSuccess,
  shifterConfigurationUpdateError,
};
