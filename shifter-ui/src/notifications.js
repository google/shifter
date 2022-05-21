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

// optionally export a default object
export default {
  shifterConversionSuccess,
  notifyAxiosError,
};
