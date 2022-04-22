package gitlab_demo


deny[msg]{  
    input.jwt.claims.iss != "gitlab.com"
    msg = "issuer is not gitlab.com"
}