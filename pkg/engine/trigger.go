package engine

// Trigger represents the first step of a workflow that watches for events before running subsequent steps.
type Trigger struct {
	Sub Subscription
}

// Subscription represent a subscription type that can send payload to its subscribers.
type Subscription interface {
}

// RESTHook is a webhook with subscription management ability.
type RESTHook struct {
	CallbackURL  string
	ClientDomain string
}

// Poller polls for information on a particular endpoint periodically.
type Poller struct {
	Interval int
}
