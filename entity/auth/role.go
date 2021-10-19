package auth

import (
	"errors"
	"strings"
)

var (
	ErrPermissionDenied = errors.New("permission denied")
)

type Role struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Accesses    []string `json:"accesses"`
}

// IsAllowed if Role's permissions matched the given permission return nil, otherwise ErrPermissionDenied
// time-complex = O(n*m)
// n = Accesses length
// m = permission's depth e.g. "auth.user.self.write" depth = 4
func (r Role) IsAllowed(perm Permission) error {
	// Simple accesses
	// Linear time-complex
	for _, a := range r.Accesses {
		// allowed all ("*")
		if a == MatchAll {
			return nil
		}

		// Check if role have exactly matched permission
		if a == perm {
			return nil
		}
	}

	// Check for complex permissions
	// time-complex = O(n*m)
	// E.g. "auth.user.company.#", "auth.user.#.write", "auth.*"
	permPath := strings.Split(perm, ".")
	for _, p := range r.Accesses {
		accessPath := strings.Split(p, ".")
		apLen := len(accessPath)
		for i, a := range accessPath {
			//  "auth.*"
			if a == MatchAll {
				return nil
			}

			// if permission path not match and not "#", calc next permission
			if a != permPath[i] && a != MatchWildcard {
				break
			}

			// match every path
			if i == apLen-1 {
				return nil
			}
		}
	}

	return ErrPermissionDenied
}
