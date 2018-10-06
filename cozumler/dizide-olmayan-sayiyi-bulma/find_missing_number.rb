def find_missing_number(array, n)
  total = (n+1) * (n+2) / 2
  array.each do |a|
    total -= a
  end

  total
end
