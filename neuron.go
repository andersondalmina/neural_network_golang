package main

import "math/rand"

type Predictor interface {
	Predict([]float64) float64
}

type Neuron struct {
	Weights []float64
	Bias    float64
}

func NewNeuron(inputs int, src rand.Source) *Neuron {
	r := rand.New(src)
	n := &Neuron{
		Weights: make([]float64, inputs),
		Bias:    r.Float64(),
	}

	for i := range n.Weights {
		n.Weights[i] = (r.Float64()*2 - 1) * 1000
	}

	return n
}

func (n *Neuron) Predict(inputs []float64) float64 {
	v := 0.0
	for i, input := range inputs {
		v += input * n.Weights[i]
	}

	return v + n.Bias
}
