/**
 * Package middleware
 * @file      : url_path.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 14:44
 **/

package middleware

import "github.com/coverai/api/internal/common/util"

var (
	// IgnoreAuthorizationUriPaths Ignore authorization uri
	IgnoreAuthorizationUriPaths = []string{
		"/users/signup",
	}

	// AllowedAllPathsByParameters Accessible URI
	AllowedAllPathsByParameters = []string{}

	// DetailPathMapByGet Detail URL
	DetailPathMapByGet = map[string]string{}

	// DetailPathMapByPost Detail URL
	DetailPathMapByPost = map[string]string{}

	// DetailPathMapByHeadOrPatch Detail URL
	DetailPathMapByHeadOrPatch = map[string]string{}

	// AllowedOnlyPostJsonPaths Accessible URI BY Post JSON
	AllowedOnlyPostJsonPaths = []string{}

	// AllowedAllPaths Accessible URI
	AllowedAllPaths = []string{
		"/initialize",
		"/users/signup",
	}
)

var allowedPathsMap map[string]bool

func init() {
	allowedPathsMap = make(map[string]bool)
	versions := []string{
		util.VersionPath,
		util.Version2Path,
	}

	for _, allowedPath := range AllowedAllPaths {
		for _, version := range versions {
			fullPath := GetPath(version, allowedPath)
			allowedPathsMap[fullPath] = true
		}
	}
}
