def reverse(string)
  unless string.is_a? String
    return false
  end
  splited_string = string.split("")
  reversed = []
  splited_string.reverse_each do |ss|
    reversed.push(ss)
  end
  reversed = reversed.join("")

  reversed
end

puts reverse("Muvaffakiyetsizleştiricileştiriveremeyebileceklerimizdenmişsinizcesine")