package lib

import (
	"github.com/google/uuid"
	"strings"
)

func OrderedUuidV1() (string, error) {
	var idv1 uuid.UUID
	idv1, err := uuid.NewUUID()
	return reorderUuidV1(idv1.String()), err
}

func reorderUuidV1(idv1 string) string {
	var sb strings.Builder
	sb.WriteString(idv1[14:18])
	sb.WriteString(idv1[9:13])
	sb.WriteString(idv1[0:8])
	sb.WriteString(idv1[19:23])
	sb.WriteString(idv1[24:])
	return sb.String()
}


