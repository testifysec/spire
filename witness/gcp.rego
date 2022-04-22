package gcp_demo

deny[msg]{
    input.project_id != "324322"
    msg = "Project ID is not correct"

}

deny[msg]{
    input.cluster_location != "us-east1-b"
    msg = "Cluster location is not correct"
}

deny[msg]{
    input.jwt.claims.iss != "https://accounts.google.com"
    msg = "JWT issuer is not correct"
}

