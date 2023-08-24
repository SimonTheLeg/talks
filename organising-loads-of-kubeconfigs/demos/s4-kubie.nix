{ pkgs ? import <nixpkgs> { } }:
pkgs.mkShell {

  packages = [ pkgs.fzf pkgs.kubie pkgs.kubectl pkgs.starship pkgs.ps ];

  # Note: For keeping this demo contained, I alias kubie to use the kubeconfig directory.
  # For non-nix use-cases, you can just configure ~/.kube/kubie.yaml to point to your kubeconfigs.
  shellHook = ''
    alias clear='/usr/bin/clear'
    eval "$(starship init bash)"
    alias kc='kubectl'
    kubie () {
      command kubie "$@" --kubeconfig="./kubeconfigs/all-in-one/config.yaml"
    }
  '';
}
