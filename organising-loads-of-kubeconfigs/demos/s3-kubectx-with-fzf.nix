{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {

  packages = [ pkgs.fzf pkgs.kubectx pkgs.kubectl pkgs.starship ];

  shellHook = ''
    alias clear='/usr/bin/clear'
    export KUBECONFIG=./kubeconfigs/all-in-one/config.yaml
    eval "$(starship init bash)"
    alias kc='kubectl'
  '';
}
