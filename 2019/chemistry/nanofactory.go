package chemistry

import (
	"github.com/mfesenko/adventofcode/2019/math"
)

// Nanofactory can transform a raw material to a product by running chemical reactions
type Nanofactory struct {
	reactions   map[string]Reaction
	rawMaterial string
	product     string
}

// NewNanofactory creates a new nanofactory
func NewNanofactory(reactions map[string]Reaction, rawMaterial string, product string) *Nanofactory {
	return &Nanofactory{
		reactions:   reactions,
		rawMaterial: rawMaterial,
		product:     product,
	}
}

// Produce calculates the amount of product that can be produced from given amount of raw material
func (n *Nanofactory) Produce(rawMaterialAmount int64) int64 {
	totalProductAmount := int64(0)
	leftovers := newFlask()
	productCost := n.ProductCost()
	for productAmount := rawMaterialAmount / productCost; productAmount != 0; {
		rawMaterialAmount -= n.produce(productAmount, leftovers)
		totalProductAmount += productAmount
		productAmount = rawMaterialAmount / productCost
	}
	return totalProductAmount
}

// ProductCost calculates the amount of raw material needed to produce one product
func (n *Nanofactory) ProductCost() int64 {
	return n.produce(1, newFlask())
}

func (n *Nanofactory) produce(requiredAmount int64, leftovers flask) int64 {
	flask := newFlask()
	flask.Add(n.product, requiredAmount)

	for !flask.ContainsOnly(n.rawMaterial) {
		nextFlask := newFlask()
		for chemical, amount := range flask {
			if chemical == n.rawMaterial {
				nextFlask.Add(chemical, amount)
				continue
			}

			if leftoverAmount, ok := leftovers[chemical]; ok {
				usedAmount := math.Min(amount, leftoverAmount)
				amount -= usedAmount
				leftovers.Add(chemical, -usedAmount)
			}

			reaction := n.reactions[chemical]

			producedAmount := reaction.Output.Amount
			count := amount / producedAmount
			if amount%producedAmount > 0 {
				count++
			}

			notUsed := (count * producedAmount) - amount
			leftovers.Add(chemical, notUsed)

			for _, input := range reaction.Input {
				nextFlask.Add(input.Chemical, count*input.Amount)
			}
		}

		flask = nextFlask
	}

	return flask[n.rawMaterial]
}
