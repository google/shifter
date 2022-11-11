resource "kubernetes_service_v1" "svc" {
  metadata {
    name = "hello-python-svc"
    namespace = "default"
  }
  spec {
    selector = {
      app = "hello-python"
    }
    session_affinity = "ClientIP"
    port {
      name        = "service"
      port        = 5000
      target_port = 5000
      protocol    = "TCP"
    }
    type = "LoadBalancer"
  }
}
