{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {

  packages = [ pkgs.fzf pkgs.kubectl pkgs.starship ];

  # due to some weird reason nix-shell does not support parameter expansion in funcs,
  # so I have simplified this here slightly, the code in the presentation will also work in normal bash
  # also for the demo the depth is 1 to make the results a bit more organized
  shellHook = ''
    alias clear='/usr/bin/clear'
    eval "$(starship init bash)"
    alias kc='kubectl'
    kcsw() {
      kubeconf=$(find ./kubeconfigs -maxdepth 1 -type f | fzf)
      export KUBECONFIG=$kubeconf
    }
  '';
}
