# Create a local variable for the load balancer name.

output "endpoint" {
  value = "https://${resource.kubernetes_service_v1.svc.status.0.load_balancer.0.ingress.0.ip}:5000"

}
