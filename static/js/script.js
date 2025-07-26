document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('bmiForm');
    const weightInput = document.getElementById('weight');
    const heightInput = document.getElementById('height');
    
    form.addEventListener('submit', function(e) {
        const weight = parseFloat(weightInput.value);
        const height = parseFloat(heightInput.value);
        
        if (weight <= 0 || isNaN(weight)) {
            e.preventDefault();
            alert('正しい体重を入力してください');
            weightInput.focus();
            return;
        }
        
        if (height <= 0 || isNaN(height)) {
            e.preventDefault();
            alert('正しい身長を入力してください');
            heightInput.focus();
            return;
        }
        
        if (weight > 500) {
            e.preventDefault();
            if (!confirm('体重が' + weight + 'kgですか？続行しますか？')) {
                weightInput.focus();
                return;
            }
        }
        
        if (height > 300) {
            e.preventDefault();
            if (!confirm('身長が' + height + 'cmですか？続行しますか？')) {
                heightInput.focus();
                return;
            }
        }
    });
    
    [weightInput, heightInput].forEach(input => {
        input.addEventListener('input', function() {
            if (this.value < 0) {
                this.value = '';
            }
        });
    });
});