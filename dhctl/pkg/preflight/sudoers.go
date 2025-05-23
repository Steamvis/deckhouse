// Copyright 2023 Flant JSC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package preflight

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	"github.com/deckhouse/deckhouse/dhctl/pkg/app"
	"github.com/deckhouse/deckhouse/dhctl/pkg/log"
	"github.com/deckhouse/deckhouse/dhctl/pkg/system/node"
)

func (pc *Checker) CheckSudoIsAllowedForUser(ctx context.Context) error {
	if app.PreflightSkipSudoIsAllowedForUserCheck {
		log.DebugLn("sudoers preflight check is skipped")
		return nil
	}

	return callSudo(ctx, pc.nodeInterface)

}

func callSudo(ctx context.Context, nodeInterface node.Interface) error {
	cmd := nodeInterface.Command("echo")
	cmd.Sudo(ctx)

	err := cmd.Run(ctx)
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && exitErr.ExitCode() != 255 {
			return errors.New("Provided SSH user is not allowed to sudo, please check that your password is correct and that this user is in the sudoers file.")
		}
		return fmt.Errorf("Unexpected error when checking sudoers permissions for SSH user: %v", err)
	}

	return nil
}
