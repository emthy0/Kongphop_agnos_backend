package password

type StrongPasswordStepRequest struct {
  InitPassword string `json:"init_password"`
}

type StrongPasswordStepResponse struct {
  NumSteps int `json:"num_of_steps"`
}