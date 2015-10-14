require "language/go"

class MillipedeGo < Formula
  desc "Print a beautiful millipede"
  homepage "https://github.com/getmillipede/millipede-go"
  url "https://github.com/getmillipede/millipede-go/archive/v1.2.0.tar.gz"
  sha256 "90c7bd537bbd71af262bab57b2465005933915b683e9739a21957832746f5f31"

  head "https://github.com/getmillipede/millipede-go.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["CGO_ENABLED"] = "0"
    ENV.prepend_create_path "PATH", buildpath/"bin"

    mkdir_p buildpath/"src/github.com/getmillipede"
    ln_s buildpath, buildpath/"src/github.com/getmillipede/millipede-go"
    Language::Go.stage_deps resources, buildpath/"src"

    system "go", "build", "-o", "millipede-go", "./cmd/millipede-go"
    bin.install "millipede-go"

    # FIXME: add autocompletion
  end

  test do
    output = shell_output(bin/"millipede-go --version")
    assert output.include? "millipede-go version"
  end
end
