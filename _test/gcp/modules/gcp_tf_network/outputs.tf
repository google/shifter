output "vpc_network" {
  value = google_compute_network.vpc.*
}

output "subnetwork" {
  value = google_compute_subnetwork.subnet.*
}

output "nataddr" {
  value = google_compute_address.nat_gw_address.*.address
}
