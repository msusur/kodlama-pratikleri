require "minitest/autorun"
require_relative "find_missing_number"

class SimpleTest < Minitest::Unit::TestCase
  def test_find_missing_number
    array = [1,2,4,5,6]

    assert_equal find_missing_number(array, array.length), 3
  end
end