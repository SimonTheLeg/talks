{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {

  packages = [ pkgs.fzf pkgs.kubie pkgs.kubectl pkgs.starship pkgs.ps ];

  # Note: for kubie to work you need to create a ~/.kube/kubie.yaml file with the following contents
  # unfortunately kubie does not allow to parse in a custom path for this one
  # configs:
  #   include:
  #     - ./kubeconfigs/*.yaml
  #     - ./kubeconfigs/*.yml
  shellHook = ''
    alias clear='/usr/bin/clear'
    eval "$(starship init bash)"
    alias kc='kubectl'
  '';
}
