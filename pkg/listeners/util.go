package listeners

import "github.com/adibrastegarnia/asn2proto/pkg/models"

func getLastAssignment(assignments []models.Assignment) models.Assignment {
	lenAssignments := len(assignments)
	lastAssignment := assignments[lenAssignments-1]
	return lastAssignment
}
