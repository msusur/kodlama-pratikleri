require "minitest/autorun"
require_relative "reverse"

class SimpleTest < Minitest::Unit::TestCase
  def test_reverse
    assert_equal reverse("abc"), "cba"
  end
end