resource "kubernetes_deployment_v1" "example" {
  metadata {
    name = "hello-python"
    labels = {
      app = "hello-python"
    }
  }

  spec {
    replicas = 3
    selector {
      match_labels = {
        app = "hello-python"
      }
    }

    template {
      metadata {
        labels = {
          app = "hello-python"
          version = ""
        }
      }

      spec {
        container {
          image = "parasmamgain/hello-python"
          name  = "hello-python"
          image_pull_policy = "IfNotPresent"
          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }

          liveness_probe {
            http_get {
              path = "/"
              port = 80
            }

            initial_delay_seconds = 3
            period_seconds        = 3
          }
        }
      }
    }
  }
}
