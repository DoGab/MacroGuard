export function useCamera() {
  let stream = $state<MediaStream | null>(null);
  let error = $state<string | null>(null);

  async function start(options: MediaStreamConstraints = { video: { facingMode: "environment" } }) {
    try {
      error = null;
      stream = await navigator.mediaDevices.getUserMedia(options);
    } catch (err) {
      error = "Could not access camera. Please check permissions.";
      console.error("Camera error:", err);
    }
  }

  function stop() {
    if (stream) {
      stream.getTracks().forEach((track) => track.stop());
      stream = null;
    }
  }

  return {
    get stream() {
      return stream;
    },
    get error() {
      return error;
    },
    start,
    stop
  };
}
